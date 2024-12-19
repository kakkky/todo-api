package auth

import (
	"crypto/rsa"
	"crypto/x509"
	_ "embed"
	"encoding/pem"

	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kakkky/app/domain/errors"
)

type JWTAuthenticator struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

//go:embed certificate/private.pem
var rawPrivateKey []byte

//go:embed certificate/public.pem
var rawPublicKey []byte

func NewJWTAuthenticator() *JWTAuthenticator {
	// *rsa.PrivateKeyにパース
	privateKey, err := parsePrivateKey(rawPrivateKey)
	if err != nil {
		log.Fatalf("private key parse error :%v", err)
	}
	publicKey, err := parsePublicKey(rawPublicKey)
	if err != nil {
		log.Fatalf("public key parse error :%v", err)
	}
	return &JWTAuthenticator{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func parsePrivateKey(pemData []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemData)
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("unknown key type")
	}
	return privateKey, nil
}
func parsePublicKey(pemData []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pemData)
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey, _ := key.(*rsa.PublicKey)
	return publicKey, nil
}

func (ja *JWTAuthenticator) GetPublicKey() *rsa.PublicKey {
	return ja.publicKey
}

func (ja *JWTAuthenticator) GenerateToken(sub, jwtId string) *jwt.Token {
	claims := jwt.StandardClaims{
		Id:        jwtId,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   sub,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token
}

func (ja *JWTAuthenticator) SignToken(token *jwt.Token) (string, error) {
	signedToken, err := token.SignedString(ja.privateKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// 署名済みのトークンを公開鍵によって解析する
func (ja *JWTAuthenticator) VerifyToken(signedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(signedToken, func(t *jwt.Token) (interface{}, error) {
		// トークンの署名アルゴリズムをチェック
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return ja.publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}
	return token, nil
}

func (ja *JWTAuthenticator) GetJWTIDFromClaim(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}
	return claims["jti"].(string), nil
}
func (ja *JWTAuthenticator) GetSubFromClaim(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}
	return claims["sub"].(string), nil
}
func (ja *JWTAuthenticator) VerifyExpiresAt(token *jwt.Token) error {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid token claims")
	}
	exp := claims["exp"].(float64)
	if time.Now().Unix() > int64(exp) {
		return errors.New("token has expired")
	}
	return nil
}
