//go:build integration

package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kakkky/app/adapter/presentation/handler/auth"
	"github.com/kakkky/app/infrastructure/api_test/testhelper"
	dbTesthelper "github.com/kakkky/app/infrastructure/db/testhelper"
	"github.com/sebdah/goldie/v2"
)

func TestAuth_Login(t *testing.T) {
	tests := []struct {
		name     string
		wantCode int
		req      auth.LoginRequest
		gfName   string
	}{
		{
			name:     "正常系",
			wantCode: http.StatusOK,
			req: auth.LoginRequest{
				Email:    "user1@test.com",
				Password: "password",
			},
			gfName: "login_nomal",
		},
		{
			name:     "準正常系:パスワードが異なる",
			wantCode: http.StatusUnauthorized,
			req: auth.LoginRequest{
				Email:    "user1@test.com",
				Password: "invalid",
			},
			gfName: "login_seminomal_password_mismatch",
		},
		{
			name:     "準正常系:指定のemailは存在しない",
			wantCode: http.StatusUnauthorized,
			req: auth.LoginRequest{
				Email:    "notfound@test.com",
				Password: "notfound",
			},
			gfName: "login_seminomal_email_notfound",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// テストデータを投入
			dbTesthelper.SetupFixtures(t, "../testdata/fixtures/users.yml")
			// リクエストボディをマーシャル（→json）
			b, _ := json.Marshal(tt.req)
			r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(b))
			rw := httptest.NewRecorder()
			mux.ServeHTTP(rw, r)
			// ステータスコードを検証
			if rw.Code != tt.wantCode {
				t.Errorf("got %d , but want Code %d", rw.Code, tt.wantCode)
			}
			resp := testhelper.FormatJSON(
				t,
				testhelper.NormalizeJWT(t, rw.Body.Bytes()),
			)
			g := goldie.New(
				t,
				goldie.WithFixtureDir("../testdata/golden/auth"),
				goldie.WithNameSuffix(".golden.json"))
			g.Assert(t, tt.gfName, resp)
		})
	}
}

func TestAuth_Logout(t *testing.T) {
	tests := []struct {
		name     string
		wantCode int
		gfName   string
		isLogin  bool
	}{
		{
			name:     "正常系",
			wantCode: http.StatusNoContent,
			gfName:   "logout_nomal",
			isLogin:  true,
		},
		{
			name:     "準正常系:ログインしていないとログアウトできない",
			wantCode: http.StatusUnauthorized,
			gfName:   "logout_seminomal_not_loggedin",
			isLogin:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbTesthelper.SetupFixtures(t, "../testdata/fixtures/users.yml")
			r := httptest.NewRequest(http.MethodDelete, "/logout", nil)
			rw := httptest.NewRecorder()
			// ログイン状態をセットアップ
			// Authorizationヘッダーを付加する
			if tt.isLogin {
				jwtToken := testhelper.LoginForTest(t, "1")
				r.Header.Set("Authorization", "Bearer "+jwtToken)
			}
			// リクエストを送信
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
				rw.Body.Bytes(),
			)
			g := goldie.New(
				t,
				goldie.WithFixtureDir("../testdata/golden/auth"),
				goldie.WithNameSuffix(".golden.json"))
			g.Assert(t, tt.gfName, resp)
		})
	}
}
