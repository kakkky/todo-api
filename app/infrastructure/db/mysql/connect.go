package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kakkky/app/config"
)

// パッケージ変数としてDB接続を管理
var (
	db *sql.DB
)

// パッケージ変数として*sql.DBをセット
func setDB(d *sql.DB) {
	db = d
}

func GetDB() *sql.DB {
	return db
}

func NewDB(ctx context.Context, cfg *config.Config) func() {
	db, close, err := connect(
		ctx,
		cfg.MySQL.User,
		cfg.MySQL.Password,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.Name,
	)
	if err != nil {
		// DB接続失敗は復旧不可
		panic(err)
	}
	// パッケージ変数にセット
	setDB(db)
	return close
}

const (
	maxRetriesCount = 5
	delay           = 5 * time.Second
)

// DBへ接続
// 失敗したらリトライさせる
func connect(ctx context.Context, user, password, host, port, name string) (*sql.DB, func(), error) {
	for i := 0; i < maxRetriesCount; i++ {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, name)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to open db : %w", err)
		}
		// 接続確認
		if err := db.PingContext(ctx); err == nil {
			// 呼び出しもと（main）でクローズ処理を強制させるために関数を返す
			return db, func() { db.Close() }, nil
		}
		// 接続できなかったらリトライ
		log.Printf("could not connect to db: %v", err)
		log.Printf("retrying in %v seconds...", delay/time.Second)
		time.Sleep(delay)
	}
	return nil, nil, fmt.Errorf("could not connect to db after %d attempts", maxRetriesCount)
}
