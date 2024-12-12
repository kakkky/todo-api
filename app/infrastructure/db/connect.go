package db

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

// DBをに接続する
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
	// dbパッケージ変数にセット
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
			// DB接続の初期エラーが発生した場合のログ
			log.Printf("failed to open db (attempt %d/%d): %v", i+1, maxRetriesCount, err)
		} else {
			// 接続確認
			if err := db.PingContext(ctx); err == nil {
				// 成功した場合、呼び出し元でdb.Close()を実行できるようにする
				log.Printf("successfully connected to db (attempt %d/%d)", i+1, maxRetriesCount)
				return db, func() { db.Close() }, nil
			}
			// 接続できた場合のエラーメッセージ
			log.Printf("failed to ping db: %v", err)
		}

		// 接続できなかったらリトライ
		log.Printf("could not connect to db (attempt %d/%d), retrying in %v seconds...", i+1, maxRetriesCount, delay/time.Second)
		time.Sleep(delay)
	}

	// 最大試行回数を超えても接続できなかった場合
	return nil, nil, fmt.Errorf("could not connect to db after %d attempts", maxRetriesCount)
}
