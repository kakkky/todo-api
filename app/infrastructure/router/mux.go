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
		mux.Handle("GET /health", composeMiddlewares(cors, logger)(http.HandlerFunc(health.HealthCheckHandler)))
		mux.Handle("GET /swagger/", swagger.WrapHandler)
	}
	// 認証系ルーティング
	{
		mux.Handle("POST /login", composeMiddlewares(cors, logger)(loginHandler))
		mux.Handle("DELETE /logout", composeMiddlewares(cors, logger, authorization)(logoutHandler))
	}
	// ユーザー系ルーティング
	{
		mux.Handle("POST /users", composeMiddlewares(cors, logger)(postUserHandler))
		mux.Handle("DELETE /users/me", composeMiddlewares(cors, logger, authorization)(deleteUserHandler))
		mux.Handle("GET /users", composeMiddlewares(cors, logger, authorization)(getUsersHandler))
		mux.Handle("GET /users/me", composeMiddlewares(cors, logger, authorization)(getCurrentUserHandler))
		mux.Handle("PATCH /users/me", composeMiddlewares(cors, logger, authorization)(updateUserHandler))
	}
	// タスク系ルーティング
	{
		mux.Handle("POST /tasks", composeMiddlewares(cors, logger, authorization)(postTaskHandler))
		mux.Handle("DELETE /tasks/{id}", composeMiddlewares(cors, logger, authorization)(deleteTaskHandler))
		mux.Handle("PATCH /tasks/{id}/state", composeMiddlewares(cors, logger, authorization)(updateTaskStateHandler))
		mux.Handle("GET /tasks/{id}", composeMiddlewares(cors, logger, authorization)(getTaskHandler))
		mux.Handle("GET /tasks", composeMiddlewares(cors, logger, authorization)(getTasksHandler))
		mux.Handle("GET /users/me/tasks", composeMiddlewares(cors, logger, authorization)(getUserTasksHandler))
	}
}
