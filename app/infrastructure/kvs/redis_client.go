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
	// Redisクライアントを作成
	redisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		// contextによるタイムアウトを許可
		ContextTimeoutEnabled: true,
	})

	// Redisが接続できるか確認するため、pingコマンドを実行
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	} else {
		log.Println("successfully connected to redis")
	}

	return redisClient
}
