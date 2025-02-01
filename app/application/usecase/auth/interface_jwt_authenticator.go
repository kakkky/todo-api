package auth

// トークンを生成する処理するインターフェース
//
//go:generate mockgen -package=auth -source=./interface_jwt_authenticator.go -destination=./mock_jwt_authenticator.go
type JwtAuthenticator interface {
	GenerateJwtToken(sub, jti string) (string, error)
	VerifyJwtToken(jwtToken string) (sub string, jti string, err error)
}
