//go:build integration

package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kakkky/app/adapter/presentation/handler/user"
	"github.com/kakkky/app/infrastructure/api_test/testhelper"
	dbTesthelper "github.com/kakkky/app/infrastructure/db/testhelper"
	"github.com/sebdah/goldie/v2"
)

func TestUser_GetUsers(t *testing.T) {
	tests := []struct {
		name     string
		wantCode int
		gfName   string
		isLogin  bool //ログインさせるか
	}{
		{
			name:     "正常系",
			wantCode: http.StatusOK,
			gfName:   "get_users_nomal",
			isLogin:  true,
		},
		{
			name:     "準正常系：未ログイン",
			wantCode: http.StatusUnauthorized,
			gfName:   "get_users_seminomal_not_loggedin",
			isLogin:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbTesthelper.SetupFixtures("../testdata/fixtures/users.yml")

			// リクエストボディをマーシャル（→json）
			r := httptest.NewRequest(http.MethodGet, "/users", nil)
			rw := httptest.NewRecorder()
			// ログイン状態をセットアップ
			// Authorizationヘッダーを付加する
			if tt.isLogin {
				jwtToken := testhelper.SetupLogin("1")
				defer testhelper.CleanupLogin("1")
				r.Header.Set("Authorization", "Bearer "+jwtToken)
			}
			// リクエストを送信
			mux.ServeHTTP(rw, r)
			// ステータスコードを検証
			if rw.Code != tt.wantCode {
				t.Errorf("got %d , but want Code %d", rw.Code, tt.wantCode)
			}
			resp := testhelper.FormatJSON(
				t,
				rw.Body.Bytes(),
			)
			g := goldie.New(
				t,
				goldie.WithFixtureDir("../testdata/golden/user"),
				goldie.WithNameSuffix(".golden.json"))
			g.Assert(t, tt.gfName, resp)
		})
	}
}

func TestUser_PostUser(t *testing.T) {
	tests := []struct {
		name     string
		wantCode int
		req      user.PostUserRequest
		gfName   string
	}{
		{
			name:     "正常系",
			wantCode: http.StatusCreated,
			req: user.PostUserRequest{
				Email:    "user0@test.com",
				Name:     "user0",
				Password: "password",
			},
			gfName: "post_user_nomal",
		},
		{
			name:     "準正常系:すでに登録しているメールアドレス",
			wantCode: http.StatusBadRequest,
			req: user.PostUserRequest{
				Email:    "user1@test.com",
				Name:     "user1",
				Password: "password",
			},
			gfName: "post_user_seminomal_dup_email",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbTesthelper.SetupFixtures("../testdata/fixtures/users.yml")

			// リクエストボディをマーシャル（→json）
			b, _ := json.Marshal(tt.req)
			r := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(b))
			rw := httptest.NewRecorder()
			mux.ServeHTTP(rw, r)
			// ステータスコードを検証
			if rw.Code != tt.wantCode {
				t.Errorf("got %d , but want Code %d", rw.Code, tt.wantCode)
			}
			resp := testhelper.FormatJSON(
				t,
				testhelper.NormalizeULID(rw.Body.Bytes()),
			)
			g := goldie.New(
				t,
				goldie.WithFixtureDir("../testdata/golden/user"),
				goldie.WithNameSuffix(".golden.json"))
			g.Assert(t, tt.gfName, resp)
		})
	}
}

func TestUser_UpdateUser(t *testing.T) {
	tests := []struct {
		name     string
		wantCode int
		req      user.UpdateUserRequest
		gfName   string
		isLogin  bool //ログインさせるか
	}{
		{
			name:     "正常系 : id=1のユーザーを更新する",
			wantCode: http.StatusOK,
			req: user.UpdateUserRequest{
				Email: "updated@updated.com",
				Name:  "updated",
			},
			gfName:  "update_user_nomal",
			isLogin: true,
		},
		{
			name:     "準正常系 : 未ログイン",
			wantCode: http.StatusUnauthorized,
			req: user.UpdateUserRequest{
				Email: "updated@updated.com",
				Name:  "updated",
			},
			gfName:  "update_user_seminormal_not_loggedin",
			isLogin: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbTesthelper.SetupFixtures("../testdata/fixtures/users.yml")

			// リクエストボディをマーシャル（→json）
			b, _ := json.Marshal(tt.req)
			r := httptest.NewRequest(http.MethodPatch, "/user", bytes.NewBuffer(b))
			rw := httptest.NewRecorder()
			// ログイン状態をセットアップ
			// Authorizationヘッダーを付加する
			if tt.isLogin {
				jwtToken := testhelper.SetupLogin("1")
				defer testhelper.CleanupLogin("1")
				r.Header.Set("Authorization", "Bearer "+jwtToken)
			}
			// リクエストを送信
			mux.ServeHTTP(rw, r)
			// ステータスコードを検証
			if rw.Code != tt.wantCode {
				t.Errorf("got %d , but want Code %d", rw.Code, tt.wantCode)
			}
			resp := testhelper.FormatJSON(
				t,
				testhelper.NormalizeULID(rw.Body.Bytes()),
			)
			g := goldie.New(
				t,
				goldie.WithFixtureDir("../testdata/golden/user"),
				goldie.WithNameSuffix(".golden.json"))
			g.Assert(t, tt.gfName, resp)
		})
	}
}

func TestUser_DeleteUser(t *testing.T) {
	tests := []struct {
		name     string
		gfName   string
		wantCode int
		isLogin  bool
	}{
		{
			name:     "正常系:id=1のユーザーを削除",
			wantCode: http.StatusNoContent,
			gfName:   "delete_user_nomal",
			isLogin:  true,
		},
		{
			name:     "準正常系:未ログイン",
			wantCode: http.StatusUnauthorized,
			gfName:   "delete_user_seminomal_not_loggedin",
			isLogin:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbTesthelper.SetupFixtures("../testdata/fixtures/users.yml")

			r := httptest.NewRequest(http.MethodDelete, "/user", nil)
			rw := httptest.NewRecorder()
			// ログイン状態をセットアップ
			// Authorizationヘッダーを付加する
			if tt.isLogin {
				jwtToken := testhelper.SetupLogin("1")
				defer testhelper.CleanupLogin("1")
				r.Header.Set("Authorization", "Bearer "+jwtToken)
			}
			mux.ServeHTTP(rw, r)
			// ステータスコードを検証
			if rw.Code != tt.wantCode {
				t.Errorf("got %d , but want Code %d", rw.Code, tt.wantCode)
			}
			// 204ステータスコードなら、JSONは返らないのでここで終了
			if rw.Code == http.StatusNoContent {
				return
			}
			resp := testhelper.FormatJSON(
				t,
				testhelper.NormalizeULID(rw.Body.Bytes()),
			)
			g := goldie.New(
				t,
				goldie.WithFixtureDir("../testdata/golden/user"),
				goldie.WithNameSuffix(".golden.json"))
			g.Assert(t, tt.gfName, resp)
		})
	}
}
