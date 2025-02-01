package testhelper

import (
	"context"
	"time"

	"github.com/kakkky/app/adapter/repository"
	"github.com/kakkky/app/infrastructure/auth"
	"github.com/kakkky/app/infrastructure/kvs"
)

func SetupLogin(id string) string {
	jwtAuthenticator := auth.NewJwtAuthenticator()
	// トークン生成
	jwtToken, _ := jwtAuthenticator.GenerateJwtToken(id, "jti")
	jwtAuthenticatorRepository := repository.NewJwtAuthenticatorRepository(kvs.NewRedisCommander())
	// Redisに保存
	jwtAuthenticatorRepository.Save(context.Background(), time.Duration(2*time.Hour), id, "jti")
	return jwtToken
}

func CleanupLogin(id string) {
	jwtAuthenticatorRepository := repository.NewJwtAuthenticatorRepository(kvs.NewRedisCommander())
	jwtAuthenticatorRepository.Delete(context.Background(), id)
}
