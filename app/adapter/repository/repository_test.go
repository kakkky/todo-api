package repository

import (
	"log"
	"testing"

	"github.com/kakkky/app/infrastructure/db/container"
)

func TestMain(m *testing.M) {
	pool, resource := container.NewDockerTestContainer()
	log.Println("success container")
	defer container.RemoveDockerTestContainer(pool, resource)

	db := container.NewDB(pool, resource)
	log.Println("success db")
	defer db.Close()
	container.SetupDB()
	log.Println("success setup")
	m.Run()
	log.Println("close....")
}
