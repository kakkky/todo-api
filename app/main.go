package main

import (
	"context"
	"log"

	"github.com/kakkky/app/config"
	_ "github.com/kakkky/app/docs"
	"github.com/kakkky/app/infrastructure/db"
	"github.com/kakkky/app/infrastructure/kvs"
	"github.com/kakkky/app/infrastructure/router"
	"github.com/kakkky/app/infrastructure/server"
)

// @title                      TodoRestAPI
// @version                    1.0
// @description                This is TodoRestPI by golang.
// @host                       localhost:8080
// @BasePath                   /
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
func main() {
	ctx := context.Background()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to read config : %v", err)
	}
	run(ctx, cfg)
}

func run(ctx context.Context, cfg *config.Config) {
	// データベース接続を初期化し、終了時にクローズする
	close := db.NewDB(ctx, cfg)
	defer close()

	// Redisクライアントに接続
	redisClient := kvs.NewRedisClient(ctx, cfg)
	defer redisClient.Close()

	// サーバーを起動し、指定したポートでリクエストを処理
	srv := server.NewServer(cfg.Server.Port, router.NewMux())
	srv.Run(ctx)
}
