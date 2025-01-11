package testhelper

import (
	"context"
	"time"

	"github.com/kakkky/app/adapter/repository"
	"github.com/kakkky/app/infrastructure/auth"
	"github.com/kakkky/app/infrastructure/kvs"
)

func SetupLogin(id string) string {
	tokenAuthenticator := auth.NewJWTAuthenticator()
	// トークン生成
	token := tokenAuthenticator.GenerateToken(id, "jti")
	tokenAuthenticatorRepository := repository.NewTokenAuthenticatorRepository(kvs.NewRedisCommander())
	// Redisに保存
	tokenAuthenticatorRepository.Save(context.Background(), time.Duration(2*time.Hour), id, "jti")
	// 署名
	signedToken, _ := tokenAuthenticator.SignToken(token)
	return signedToken
}

func CleanupLogin(id string) {
	tokenAuthenticatorRepository := repository.NewTokenAuthenticatorRepository(kvs.NewRedisCommander())
	tokenAuthenticatorRepository.Delete(context.Background(), id)
}
