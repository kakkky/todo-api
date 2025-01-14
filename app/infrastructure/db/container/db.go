package container

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kakkky/app/domain/errors"
	"github.com/ory/dockertest/v3"
)

// DBに接続する
func NewDB(pool *dockertest.Pool, resource *dockertest.Resource) *sql.DB {
	var db *sql.DB
	// エラーだと再実行を繰り返す
	if err := pool.Retry(func() error {
		var err error
		// 公開ポート番号を取得
		port, err = strconv.Atoi(resource.GetPort("3306/tcp")) //コンテナ内のポート番号は3306
		if err != nil {
			return err
		}
		// DBに接続
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", userName, password, hostname, port, dbName))
		if err != nil {
			return err
		}
		// 接続確認
		return db.Ping()
	}); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return db
}

func SetupDB() {
	migrationsPath := getMigrationsPath()
	m, err := migrate.New(migrationsPath, fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s?parseTime=true", userName, password, hostname, port, dbName))
	if err != nil {
		log.Fatalf("failed to create migrate instance:%v", err)
	}
	// マイグレーションをテストDBに適用させる
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("migrations have already been applied")
		} else {
			log.Fatalf("failed to up migrations:%v", err)
		}
	}
}

const migrationsRelativePath = "../migrations"

// どこからSetupDBを呼び出しても/migrationsへのパスを取得できるようにする
func getMigrationsPath() string {
	// Callerを呼んだファイル名（ファイルパス）を取得する
	_, callerFile, _, ok := runtime.Caller(0)
	print(callerFile)
	if !ok {
		log.Fatal("failed to get caller directory")
	}
	// ファイルパスからディレクトリ部分を取り出す
	callerDir := filepath.Dir(callerFile)
	// /migrationsへの絶対パスを作成
	migrationsAbsPath := "file://" + filepath.Join(callerDir, migrationsRelativePath)
	return migrationsAbsPath
}
