package container

import (
	"log"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	userName = "root"
	password = "secret"
	hostname = "localhost"
	dbName   = "test-db"
	port     int //ポート番号は起動したコンテナから取得する（ランダム）
)

// DockerTestのコンテナを起動
func NewDockerTestContainer() (*dockertest.Pool, *dockertest.Resource) {
	// デフォルトでUnixソケットを使用する
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("failled to construct pool:%v", err)
	}
	// Dockerに接続確認
	if err := pool.Client.Ping(); err != nil {
		log.Fatalf("failed to connect to Docker: %v", err)
	}
	// コンテナの起動設定を指定
	runOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=" + password,
			"MYSQL_DATABASE=" + dbName,
		},
		// 開発環境の設定に揃える
		Cmd: []string{
			"mysqld",
			"--character-set-server=utf8mb4",
			"--collation-server=utf8mb4_unicode_ci",
		},
	}
	resource, err := pool.RunWithOptions(runOptions, func(hc *docker.HostConfig) {
		// コンテナは停止後に自動削除される
		hc.AutoRemove = true
		// コンテナが異常終了しても再起動しない
		hc.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("failed to start resource: %v", err)
	}
	return pool, resource
}

// コンテナを削除する
func RemoveDockerTestContainer(pool *dockertest.Pool, resource *dockertest.Resource) {
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("failed to  purge resource: %s", err)
	}
}
