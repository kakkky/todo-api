package repository

import (
	"context"

	"time"

	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/app/infrastructure/kvs"
	"github.com/redis/go-redis/v9"
)

type tokenAuthenticatorRepository struct{}

func NewTokenAuthenticatorRepository() *tokenAuthenticatorRepository {
	return &tokenAuthenticatorRepository{}
}

func (tar *tokenAuthenticatorRepository) Save(ctx context.Context, duration time.Duration, userID, jwtID string) error {
	cli := kvs.GetRedisClient()

	// Redis クライアントが nil の場合はエラーを返す
	if cli == nil {
		return errors.New("failed to get Redis client")
	}
	status := cli.Set(ctx, userID, jwtID, duration)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

// 存在しないKEYを指定した場合は空文字を返す
func (tar *tokenAuthenticatorRepository) Load(ctx context.Context, userID string) (string, error) {
	cli := kvs.GetRedisClient()
	status := cli.Get(ctx, userID)
	if status.Err() != nil {
		// nilだったら空文字を返す
		if status.Err() == redis.Nil {
			return "", nil
		}
		return "", status.Err()
	}
	return status.Val(), nil
}

func (tar *tokenAuthenticatorRepository) Delete(ctx context.Context, userID string) error {
	cli := kvs.GetRedisClient()
	status := cli.Del(ctx, userID)
	if status.Err() != nil {
		if status.Err() == redis.Nil {
			return nil
		}
		return status.Err()
	}
	return nil
}
