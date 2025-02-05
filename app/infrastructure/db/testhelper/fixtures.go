package testhelper

import (
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/kakkky/app/infrastructure/db"
)

func SetupFixtures(t *testing.T, fixture_path ...string) {
	t.Helper()

	fixtures, err := testfixtures.New(
		testfixtures.Database(db.GetDB()),
		testfixtures.Dialect("mysql"),
		testfixtures.SkipResetSequences(),
		testfixtures.Files(
			fixture_path...,
		),
	)
	if err != nil {
		t.Fatalf("testfixtures failed to create Loader instance:%v", err)
	}
	// テーブルのデータを削除&用意
	if err := fixtures.Load(); err != nil {
		t.Fatalf("testfixtures failed to load fixtures:%v", err)
	}
}
