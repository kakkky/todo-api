package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kakkky/app/domain/errors"
	"golang.org/x/sync/errgroup"
)

// *http.Server型をラップする
type server struct {
	srv *http.Server
}

// 独自のserver.Server型を返すコンストラクタ
// main関数でマルチプレクサを渡す
func NewServer(port string, mux http.Handler) *server {
	return &server{
		&http.Server{
			Addr:    port,
			Handler: mux,
		},
	}
}

// サーバーを起動する
func (s *server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	// シグナルの監視をやめてリソースを開放する
	defer stop()
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		// サーバー起動中にエラーが起きると、ctxに伝達される（Shutdownによるエラーは正常なので無視）
		log.Printf("server is runnning at port %q...", s.srv.Addr)
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("http server on %s failed : %+v", s.srv.Addr, err)
			return err
		}
		log.Printf("The server on %s is gracefully shutting down", s.srv.Addr)
		return nil
	})
	// サーバーからのエラーとシグナルを待機する
	<-ctx.Done()
	// シャットダウンをするまでのタイムアウト時間を設定
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown http server on %s : %+v", s.srv.Addr, err)
	}
	log.Println("server was successfully shut down gracefully")
	// ゴルーチンのクロージャ内のエラーを返す
	// ゴルーチンを待機する
	return eg.Wait()
}
