package kvs

import (
	"log"

	"github.com/alicebob/miniredis"
	"github.com/redis/go-redis/v9"
)

func NewRedisTestClient() *redis.Client {
	// redisサーバを作る
	s, err := miniredis.Run()
	if err != nil {
		log.Fatalf("failed to run miniredis :%v", err)
	}
	// *redis.Clientをパッケージ変数にセット
	redisClient = redis.NewClient(&redis.Options{
		Addr:     s.Addr(),
		Password: "",
		DB:       0,
	})
	return redisClient
}
