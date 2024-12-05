package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	// テストとしてキャンセル通知を送るため
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	// サーバー構造体（テスト対象）
	mux := http.NewServeMux()
	sut := NewServer(":8081", mux)
	// デフォルトマルチプレクサにハンドラを登録
	want := "Hello,world!"
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, want)
	})
	eg.Go(func() error {
		return sut.Run(ctx)
	})

	// リクエストを送ってレスポンスが期待通りか確かめる
	resp, err := http.Get("http://localhost:8081")
	if err != nil {
		t.Errorf("failed to request server : %v", err)
	}
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}
	defer resp.Body.Close()
	if string(got) != want {
		t.Errorf("want %q , but got %q", want, got)
	}
	// キャンセル通知を送ってサーバーが正常に終了するか
	cancel()
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
