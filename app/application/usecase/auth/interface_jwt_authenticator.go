package auth

// トークンを生成する処理するインターフェース
//
//go:generate mockgen -package=auth -source=./interface_jwt_authenticator.go -destination=./mock_jwt_authenticator.go
type JwtAuthenticator interface {
	GenerateJwtToken(sub, jti string) (string, error)
	// SignToken(token *jwt.Token) (string, error)
	VerifyJwtToken(signedToken string) (sub string, jti string, err error)
	// GetJtiFromClaim(token *jwt.Token) (string, error)
	// GetSubFromClaim(token *jwt.Token) (string, error)
	// VerifyExpiresAt(token *jwt.Token) error
}
