package router

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/handler/health"
	swagger "github.com/swaggo/http-swagger"
)

// ルーティングを登録したマルチプレクサを返す
func NewMux() http.Handler {
	// マルチプレクサを初期化
	mux := http.NewServeMux()
	// ハンドラーの初期化
	initHandlers()
	// ミドルウェアの初期化
	initMiddlewares()
	// ハンドラをルーティングに登録
	registerRoutes(mux)

	return mux
}

func registerRoutes(mux *http.ServeMux) {
	// 開発用ルーティング
	{
		mux.HandleFunc("GET /health", health.HealthCheckHandler)
		mux.Handle("GET /swagger/", swagger.WrapHandler)
	}
	// 認証系ルーティング
	{
		mux.Handle("POST /login", composeMiddlewares(logger)(loginHandler))
		mux.Handle("DELETE /logout", composeMiddlewares(logger, authorization)(logoutHandler))
	}
	// ユーザー系ルーティング
	{
		mux.Handle("POST /users", composeMiddlewares(logger)(postUserHandler))
		mux.Handle("DELETE /users/me", composeMiddlewares(logger, authorization)(deleteUserHandler))
		mux.Handle("GET /users", composeMiddlewares(logger, authorization)(getUsersHandler))
		mux.Handle("PATCH /users/me", composeMiddlewares(logger, authorization)(updateUserHandler))
	}
	// タスク系ルーティング
	{
		mux.Handle("POST /tasks", composeMiddlewares(logger, authorization)(postTaskHandler))
		mux.Handle("DELETE /tasks/{id}", composeMiddlewares(logger, authorization)(deleteTaskHandler))
		mux.Handle("PATCH /tasks/{id}/state", composeMiddlewares(logger, authorization)(updateTaskStateHandler))
		mux.Handle("GET /tasks/{id}", composeMiddlewares(logger, authorization)(getTaskHandler))
		mux.Handle("GET /tasks", composeMiddlewares(logger, authorization)(getTasksHandler))
		mux.Handle("GET /users/me/tasks", composeMiddlewares(logger, authorization)(getUserTasksHandler))
	}
}
