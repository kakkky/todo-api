package auth

import (
	"github.com/golang-jwt/jwt"
)

// トークンを生成する処理するインターフェース
//
//go:generate mockgen -package=auth -source=./interface_token_authenticator.go -destination=./mock_token_authenticator.go
type TokenAuthenticator interface {
	GenerateToken(sub, jwtID string) *jwt.Token
	SignToken(token *jwt.Token) (string, error)
	VerifyToken(signedToken string) (*jwt.Token, error)
	GetJWTIDFromClaim(token *jwt.Token) (string, error)
	GetSubFromClaim(token *jwt.Token) (string, error)
	VerifyExpiresAt(token *jwt.Token) error
}
