package repository

import (
	"log"
	"testing"

	"github.com/kakkky/app/infrastructure/db"
	"github.com/kakkky/app/infrastructure/db/container"
	"github.com/kakkky/app/infrastructure/db/sqlc"
)

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
	m.Run()
}
