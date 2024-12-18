package kvs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kakkky/app/config"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

// パッケージ変数を取得
func GetRedisClient() *redis.Client {
	return redisClient
}

func NewRedisClient(ctx context.Context, cfg *config.Config) *redis.Client {
	// パッケージ変数にセット
	redisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		// contextによるタイムアウトを許可
		ContextTimeoutEnabled: true,
	})
	log.Println("success to connect to redis client")
	return redisClient
}
