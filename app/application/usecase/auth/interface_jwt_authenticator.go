package auth

import (
	"github.com/golang-jwt/jwt"
)

// トークンを生成する処理するインターフェース
//
//go:generate mockgen -package=auth -source=./interface_jwt_authenticator.go -destination=./mock_jwt_authenticator.go
type JwtAuthenticator interface {
	GenerateToken(sub, jwtID string) *jwt.Token
	SignToken(token *jwt.Token) (string, error)
	VerifyToken(signedToken string) (*jwt.Token, error)
	GetJwtIDFromClaim(token *jwt.Token) (string, error)
	GetSubFromClaim(token *jwt.Token) (string, error)
	VerifyExpiresAt(token *jwt.Token) error
}
