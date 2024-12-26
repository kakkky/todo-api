//go:build integration_write

package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kakkky/app/adapter/presentation/user"
	"github.com/kakkky/app/infrastructure/api_test/testhelper"
	dbTesthelper "github.com/kakkky/app/infrastructure/db/testhelper"
	"github.com/sebdah/goldie/v2"
)

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
				Email:    "user1@test.com",
				Name:     "user1",
				Password: "password",
			},
			gfName: "post_user_nomal",
		},
		{
			name:     "準正常系:すでに登録しているメールアドレス",
			wantCode: http.StatusBadRequest,
			req: user.PostUserRequest{
				Email:    "user0@test.com",
				Name:     "user0",
				Password: "password",
			},
			gfName: "post_user_seminomal_dup_email",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbTesthelper.SetupFixtures(t, "../testdata/fixtures/users.yml")
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
			name:     "正常系 : id=0のユーザーを更新する",
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
			dbTesthelper.SetupFixtures(t, "../testdata/fixtures/users.yml")

			// リクエストボディをマーシャル（→json）
			b, _ := json.Marshal(tt.req)
			r := httptest.NewRequest(http.MethodPatch, "/user", bytes.NewBuffer(b))
			rw := httptest.NewRecorder()
			// ログイン状態をセットアップ
			// Authorizationヘッダーを付加する
			if tt.isLogin {
				signedToken := testhelper.SetupLogin("0")
				defer testhelper.CleanupLogin("0")
				r.Header.Set("Authorization", "Bearer "+signedToken)
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
		wantCode int
		isLogin  bool
	}{
		{
			name:     "正常系:id=0のユーザーを削除",
			wantCode: http.StatusNoContent,
			isLogin:  true,
		},
		{
			name:     "準正常系:未ログイン",
			wantCode: http.StatusUnauthorized,
			isLogin:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbTesthelper.SetupFixtures(t, "../testdata/fixtures/users.yml")

			r := httptest.NewRequest(http.MethodDelete, "/user", nil)
			rw := httptest.NewRecorder()
			// ログイン状態をセットアップ
			// Authorizationヘッダーを付加する
			if tt.isLogin {
				signedToken := testhelper.SetupLogin("0")
				defer testhelper.CleanupLogin("0")
				r.Header.Set("Authorization", "Bearer "+signedToken)
			}
			mux.ServeHTTP(rw, r)
			// ステータスコードを検証
			if rw.Code != tt.wantCode {
				t.Errorf("got %d , but want Code %d", rw.Code, tt.wantCode)
			}
		})
	}
}
