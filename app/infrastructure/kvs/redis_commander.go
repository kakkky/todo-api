package kvs

import (
	"context"
	"errors"

	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCommander struct {
}

func NewRedisCommander() *RedisCommander {
	return &RedisCommander{}
}

func (rc *RedisCommander) Save(ctx context.Context, duration time.Duration, userID, jwtID string) error {
	cli := GetRedisClient()
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

func (rc *RedisCommander) Load(ctx context.Context, userID string) (string, error) {
	cli := GetRedisClient()
	// Redis クライアントが nil の場合はエラーを返す
	if cli == nil {
		return "", errors.New("failed to get Redis client")
	}
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

func (rc *RedisCommander) Delete(ctx context.Context, userID string) error {
	cli := GetRedisClient()
	// Redis クライアントが nil の場合はエラーを返す
	if cli == nil {
		return errors.New("failed to get Redis client")
	}
	status := cli.Del(ctx, userID)
	if status.Err() != nil {
		if status.Err() == redis.Nil {
			return nil
		}
		return status.Err()
	}
	return nil
}
