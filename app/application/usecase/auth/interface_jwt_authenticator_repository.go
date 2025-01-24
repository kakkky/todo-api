package auth

import (
	"context"
	"time"
)

// トークンをKVSで操作するリポジトリインターフェース

//go:generate mockgen -package=auth -source=./interface_jwt_authenticator_repository.go -destination=./mock_jwt_authenticator_repository.go
type JwtAuthenticatorRepository interface {
	Save(ctx context.Context, duration time.Duration, userID, jti string) error
	Load(ctx context.Context, userID string) (string, error)
	Delete(ctx context.Context, userID string) error
}
