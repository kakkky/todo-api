package repository

import (
	"context"

	"time"
)

type jwtAuthenticatorRepository struct {
	kvsCommander KvsCommander
}

func NewJwtAuthenticatorRepository(kvsCommander KvsCommander) *jwtAuthenticatorRepository {
	return &jwtAuthenticatorRepository{
		kvsCommander: kvsCommander,
	}
}

func (tar *jwtAuthenticatorRepository) Save(ctx context.Context, duration time.Duration, userID, jti string) error {
	return tar.kvsCommander.Save(ctx, duration, userID, jti)
}

// 存在しないKEYを指定した場合は空文字を返す
func (tar *jwtAuthenticatorRepository) Load(ctx context.Context, userID string) (string, error) {
	return tar.kvsCommander.Load(ctx, userID)
}

func (tar *jwtAuthenticatorRepository) Delete(ctx context.Context, userID string) error {
	return tar.kvsCommander.Delete(ctx, userID)
}
