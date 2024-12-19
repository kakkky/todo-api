package repository

import (
	"context"
	"time"

	"github.com/kakkky/app/infrastructure/kvs"
	"github.com/redis/go-redis/v9"
)

type tokenAuthenticatorRepository struct{}

func NewTokenAuthenticatorRepository() *tokenAuthenticatorRepository {
	return &tokenAuthenticatorRepository{}
}

func (tar *tokenAuthenticatorRepository) Save(ctx context.Context, duration time.Duration, userID, jwtID string) error {
	cli := kvs.GetRedisClient()
	status := cli.Set(ctx, userID, jwtID, duration)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

func (tar *tokenAuthenticatorRepository) Load(ctx context.Context, userID string) (string, error) {
	cli := kvs.GetRedisClient()
	status := cli.Get(ctx, userID)
	if status.Err() != nil {
		if status.Err() == redis.Nil {
			return "", nil
		}
		return "", status.Err()
	}
	return status.Val(), nil
}
