package auth

import (
	"crypto/rsa"
	_ "embed"

	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kakkky/app/domain/errors"
)

type jwtAuthenticator struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

//go:embed certificate/private.pem
var rawPrivateKey []byte

//go:embed certificate/public.pem
var rawPublicKey []byte

func NewJwtAuthenticator() *jwtAuthenticator {
	// *rsa.PrivateKeyにパース
	privateKey, err := parsePrivateKey(rawPrivateKey)
	if err != nil {
		log.Fatalf("private key parse error :%v", err)
	}
	publicKey, err := parsePublicKey(rawPublicKey)
	if err != nil {
		log.Fatalf("public key parse error :%v", err)
	}
	return &jwtAuthenticator{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func (ja *jwtAuthenticator) GenerateJwtToken(sub, jti string) (string, error) {
	claims := jwt.StandardClaims{
		Id:        jti,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   sub,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(ja.privateKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// 署名済みのトークンを公開鍵によって検証する
// クレームから取り出した値を返す
func (ja *jwtAuthenticator) VerifyJwtToken(signedToken string) (sub string, jti string, err error) {
	token, err := jwt.Parse(signedToken, func(t *jwt.Token) (interface{}, error) {
		// トークンの署名アルゴリズムをチェック
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return ja.publicKey, nil
	})
	if err != nil {
		return "", "", err
	}
	//　トークンは有効か
	if !token.Valid {
		return "", "", fmt.Errorf("token is invalid")
	}
	// 有効期限は過ぎていないか
	if err := verifyExpiresAt(token); err != nil {
		return "", "", err
	}
	// subを取得
	sub, err = getSubFromClaim(token)
	if err != nil {
		return "", "", err
	}
	// jtiを取得
	jti, err = getJtiFromClaim(token)
	if err != nil {
		return "", "", err
	}
	return sub, jti, nil
}

func getJtiFromClaim(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}
	return claims["jti"].(string), nil
}
func getSubFromClaim(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}
	return claims["sub"].(string), nil
}
func verifyExpiresAt(token *jwt.Token) error {
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
