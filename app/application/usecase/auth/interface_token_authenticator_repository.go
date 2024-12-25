package auth

import (
	"context"
	"time"
)

// トークンをKVSで操作するインターフェース

//go:generate mockgen -package=auth -source=./interface_token_authenticator_repository.go -destination=./mock_token_authenticator_repository.go
type TokenAuthenticatorRepository interface {
	Save(ctx context.Context, duration time.Duration, userID, jwtID string) error
	Load(ctx context.Context, user_id string) (string, error)
	Delete(ctx context.Context, user_id string) error
}
