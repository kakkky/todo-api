package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	// テストとしてキャンセル通知を送るため
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	// サーバー構造体（テスト対象）
	mux := http.NewServeMux()
	sut := NewServer(":8881", mux)
	// デフォルトマルチプレクサにハンドラを登録
	want := "Hello,world!"
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, want)
	})
	eg.Go(func() error {
		return sut.Run(ctx)
	})

	// リクエストを送ってレスポンスが期待通りか確かめる
	for i := 0; i < 5; i++ {
		resp, err := http.Get("http://localhost:8881")
		if err != nil {
			if i < 4 {
				time.Sleep(1 * time.Second) // リトライ前に少し待機
				continue
			}
			t.Errorf("failed to request server : %v", err)
			break
		}

		// ステータスコードを確認
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status code 200 but got %d", resp.StatusCode)
			continue
		}

		got, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read body: %v", err)
		}

		// 読み取ったレスポンスを閉じる
		resp.Body.Close()

		if string(got) != want {
			t.Errorf("want %q , but got %q", want, got)
		}
	}

	// キャンセル通知を送ってサーバーが正常に終了するか
	cancel()
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
