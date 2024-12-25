package testhelper

import (
	"context"
	"time"

	"github.com/kakkky/app/adapter/repository"
	"github.com/kakkky/app/infrastructure/auth"
)

func SetupLogin(id string) string {
	tokenAuthenticator := auth.NewJWTAuthenticator()
	token := tokenAuthenticator.GenerateToken(id, "jti")
	tokenAuthenticatorRepository := repository.NewTokenAuthenticatorRepository()
	tokenAuthenticatorRepository.Save(context.Background(), time.Duration(2*time.Hour), id, "jti")
	signedToken, _ := tokenAuthenticator.SignToken(token)
	return signedToken
}

func CleanupLogin(id string) {
	tokenAuthenticatorRepository := repository.NewTokenAuthenticatorRepository()
	tokenAuthenticatorRepository.Delete(context.Background(), id)
}
