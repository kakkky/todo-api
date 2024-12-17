package main

import (
	"context"
	"log"

	"github.com/kakkky/app/config"
	_ "github.com/kakkky/app/docs"
	"github.com/kakkky/app/infrastructure/db"
	"github.com/kakkky/app/infrastructure/db/sqlc"
	"github.com/kakkky/app/infrastructure/router"
	"github.com/kakkky/app/infrastructure/server"
)

// @title       TodoRestAPI
// @version     1.0
// @description This is TodoARestPI by golang.
// @host        localhost:8080
// @BasePath    /
func main() {
	ctx := context.Background()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to read config : %v", err)
	}
	if err := run(ctx, cfg); err != nil {
		log.Printf("error occured in process: %v", err)
	}
}

func run(ctx context.Context, cfg *config.Config) error {
	// データベース接続を初期化し、終了時にクローズする
	close := db.NewDB(ctx, cfg)
	defer close()

	// sqlc を使用したクエリ実行のために、sqlcパッケージにクエリオブジェクト変数を設定
	sqlc.SetQueries(db.GetDB())

	// サーバーを起動し、指定したポートでリクエストを処理
	srv := server.NewServer(cfg.Server.Port, router.NewMux())
	srv.Run(ctx)

	return nil
}
