package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/kakkky/app/adapter/presentation/presenter"
	"github.com/kakkky/app/application/usecase/user/auth"
	"github.com/kakkky/app/domain/errors"
)

type UserIDKey struct{}

func Authorication(tokenAuthenticator auth.TokenAuthenticator, tokenAuthenticatorRepository auth.TokenAuthenticatorRepository) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, err := authenticateRequest(r, tokenAuthenticator, tokenAuthenticatorRepository)
			if err != nil {
				presenter.RespondUnAuthorized(w, err.Error())
				return
			}
			// コンテキストにuserIDを含める
			ctx := context.WithValue(r.Context(), UserIDKey{}, userID)
			// 後続の処理（ハンドラ）を実行する
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func authenticateRequest(r *http.Request, tokenAuthenticator auth.TokenAuthenticator, tokenAuthenticatorRepository auth.TokenAuthenticatorRepository) (string, error) {
	// JWT トークンを取得
	signedToken, err := getSignedToken(r)
	if err != nil {
		return "", err
	}

	// トークンを検証
	token, err := tokenAuthenticator.VerifyToken(signedToken)
	if err != nil {
		return "", err
	}

	// トークンの有効期限を検証
	if err := tokenAuthenticator.VerifyExpiresAt(token); err != nil {
		return "", err
	}

	// JWT クレームから情報を取得
	jti, err := tokenAuthenticator.GetJWTIDFromClaim(token)
	if err != nil {
		return "", err
	}
	userID, err := tokenAuthenticator.GetSubFromClaim(token)
	if err != nil {
		return "", err
	}

	// KVS から保存された jti を取得
	jtiFromKVS, err := tokenAuthenticatorRepository.Load(r.Context(), userID)
	if err != nil {
		return "", err
	}

	// jti が一致しない場合はエラー
	if jti != jtiFromKVS {
		return "", errors.New("invalid JWT ID")
	}

	return userID, nil
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
