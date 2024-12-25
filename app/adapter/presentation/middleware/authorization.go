package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/kakkky/app/adapter/presentation/presenter"
	"github.com/kakkky/app/application/usecase/auth"
)

type UserIDKey struct{}

func Authorication(authorizationUsecase *auth.AuthorizationUsecase) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			signedToken, err := getSignedToken(r)
			if err != nil {
				presenter.RespondUnAuthorized(w, err.Error())
				return
			}
			input := auth.AuthorizationInputDTO{SignedToken: signedToken}
			output, err := authorizationUsecase.Run(r.Context(), input)
			if err != nil {
				presenter.RespondUnAuthorized(w, err.Error())
				return
			}
			userID := output.UserID
			// コンテキストにuserIDを含める
			ctx := context.WithValue(r.Context(), UserIDKey{}, userID)
			// 後続の処理（ハンドラ）を実行する
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getSignedToken(r *http.Request) (string, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return "", errors.New("Authorization Header is missing")
	}
	// スペースを区切りとして最大２つに分割
	parts := strings.SplitN(authorizationHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid Authorization header format")
	}
	return parts[1], nil
}
