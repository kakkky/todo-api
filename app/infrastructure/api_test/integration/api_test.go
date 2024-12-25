package integration

import (
	"log"
	"net/http"
	"testing"

	"github.com/kakkky/app/infrastructure/db"
	"github.com/kakkky/app/infrastructure/db/container"
	"github.com/kakkky/app/infrastructure/db/sqlc"
	"github.com/kakkky/app/infrastructure/kvs"
	"github.com/kakkky/app/infrastructure/router"
)

var mux http.Handler

func TestMain(m *testing.M) {
	//dockertestコンテナを起動
	pool, resource := container.NewDockertestContainer()
	log.Println("success to start dockertest container")
	defer func() {
		container.RemoveDockertestContainer(pool, resource)
		log.Println("success to remove dockertest container")
	}()
	// DBに接続
	testDB := container.NewDB(pool, resource)
	log.Println("success to connect test-db")
	defer testDB.Close()
	// マイグレーションを適用させる
	container.SetupDB()
	log.Println("success to apply migrations")
	// dbパッケージ変数にテスト用DBをセット
	db.SetDB(testDB)
	// sqlcパッケージ変数*Queriesをセット
	sqlc.SetQueries(db.GetDB())
	log.Println("dockertest & test-db settings complete")

	// テスト用のredisサーバーを起動
	cli := kvs.NewRedisTestClient()
	defer cli.Close()

	// ルーティングを初期化
	mux = router.NewMux()

	m.Run()
}
