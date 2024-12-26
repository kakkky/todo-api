//go:build integration_read

package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kakkky/app/infrastructure/api_test/testhelper"
	dbTesthelper "github.com/kakkky/app/infrastructure/db/testhelper"
	"github.com/sebdah/goldie/v2"
)

func TestUser_GetUsers(t *testing.T) {
	t.Parallel()
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
			dbTesthelper.SetupFixtures(t, "../testdata/fixtures/users.yml")

			// リクエストボディをマーシャル（→json）
			r := httptest.NewRequest(http.MethodGet, "/users", nil)
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
