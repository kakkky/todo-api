package testhelper

import (
	"context"
	"testing"
	"time"

	"github.com/kakkky/app/adapter/repository"
	"github.com/kakkky/app/infrastructure/auth"
	"github.com/kakkky/app/infrastructure/kvs"
)

func SetupLogin(t *testing.T, id string) string {
	t.Helper()

	jwtAuthenticator := auth.NewJwtAuthenticator()
	// トークン生成
	jwtToken, err := jwtAuthenticator.GenerateJwtToken(id, "jti")
	if err != nil {
		t.Fatalf("error occuerd in jwtAuthenticator.GenerateJwtToken() :%v", err)
	}
	jwtAuthenticatorRepository := repository.NewJwtAuthenticatorRepository(kvs.NewRedisCommander())
	// Redisに保存
	jwtAuthenticatorRepository.Save(context.Background(), time.Duration(2*time.Hour), id, "jti")
	return jwtToken
}

func CleanupLogin(t *testing.T, id string) {
	t.Helper()

	jwtAuthenticatorRepository := repository.NewJwtAuthenticatorRepository(kvs.NewRedisCommander())
	if err := jwtAuthenticatorRepository.Delete(context.Background(), id); err != nil {
		t.Fatalf("error occured in jwtAuthenticatorRepository.Delete() :%v", err)
	}
}
