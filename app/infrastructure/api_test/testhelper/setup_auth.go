package testhelper

import (
	"context"
	"testing"
	"time"

	"github.com/kakkky/app/adapter/repository"
	"github.com/kakkky/app/infrastructure/auth"
	"github.com/kakkky/app/infrastructure/kvs"
)

// ログイン状態をセットアップする
// userID：JTIのペアを、KVSに保存する
// 生成したJWTトークンを返す
func LoginForTest(t *testing.T, userID string) string {
	t.Helper()

	jwtAuthenticator := auth.NewJwtAuthenticator()
	// トークン生成
	jwtToken, err := jwtAuthenticator.GenerateJwtToken(userID, "jti")
	if err != nil {
		t.Fatalf("error occuerd in jwtAuthenticator.GenerateJwtToken() :%v", err)
	}
	jwtAuthenticatorRepository := repository.NewJwtAuthenticatorRepository(kvs.NewRedisCommander())
	// Redisに保存
	jwtAuthenticatorRepository.Save(context.Background(), time.Duration(2*time.Hour), userID, "jti")
	return jwtToken
}

func LogoutForTest(t *testing.T, userID string) {
	t.Helper()

	jwtAuthenticatorRepository := repository.NewJwtAuthenticatorRepository(kvs.NewRedisCommander())
	if err := jwtAuthenticatorRepository.Delete(context.Background(), userID); err != nil {
		t.Fatalf("error occured in jwtAuthenticatorRepository.Delete() :%v", err)
	}
}
