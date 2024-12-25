//go:build integration_write

package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kakkky/app/adapter/presentation/user"
	"github.com/kakkky/app/infrastructure/api_test/testutil"
	"github.com/sebdah/goldie/v2"
)

func TestUser_PostUser(t *testing.T) {
	tests := []struct {
		name     string
		wantCode int
		req      user.PostUserRequest
		resp     any
	}{
		{
			name:     "正常系",
			wantCode: http.StatusCreated,
			req: user.PostUserRequest{
				Email:    "user1@test.com",
				Name:     "user1",
				Password: "password",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// リクエストボディをマーシャル（→json）
			b, _ := json.Marshal(tt.req)
			r := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(b))
			rw := httptest.NewRecorder()
			mux.ServeHTTP(rw, r)
			// ステータスコードを検証
			if rw.Code != tt.wantCode {
				t.Errorf("got %d , but want Code %d", rw.Code, tt.wantCode)
			}
			resp := testutil.FormatJSON(
				t,
				testutil.NormalizeULID(rw.Body.Bytes()),
			)
			g := goldie.New(
				t,
				goldie.WithFixtureDir("../testdata"),
				goldie.WithNameSuffix(".golden.json"))
			g.Assert(t, t.Name(), resp)
		})
	}
}
