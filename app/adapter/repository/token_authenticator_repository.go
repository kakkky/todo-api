package repository

import (
	"context"

	"time"
)

type tokenAuthenticatorRepository struct {
	kvsCommander KvsCommander
}

func NewTokenAuthenticatorRepository(kvsCommander KvsCommander) *tokenAuthenticatorRepository {
	return &tokenAuthenticatorRepository{
		kvsCommander: kvsCommander,
	}
}

func (tar *tokenAuthenticatorRepository) Save(ctx context.Context, duration time.Duration, userID, jwtID string) error {
	return tar.kvsCommander.Save(ctx, duration, userID, jwtID)
}

// 存在しないKEYを指定した場合は空文字を返す
func (tar *tokenAuthenticatorRepository) Load(ctx context.Context, userID string) (string, error) {
	return tar.kvsCommander.Load(ctx, userID)
}

func (tar *tokenAuthenticatorRepository) Delete(ctx context.Context, userID string) error {
	return tar.kvsCommander.Delete(ctx, userID)
}
