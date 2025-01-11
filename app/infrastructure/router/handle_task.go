package router

import (
	"log"
	"net/http"

	taskHandler "github.com/kakkky/app/adapter/presentation/handler/task"
	"github.com/kakkky/app/adapter/presentation/middleware"
	queryservice "github.com/kakkky/app/adapter/query_service"
	"github.com/kakkky/app/adapter/repository"
	"github.com/kakkky/app/application/usecase/auth"
	taskUsecase "github.com/kakkky/app/application/usecase/task"
	authInfra "github.com/kakkky/app/infrastructure/auth"
	"github.com/kakkky/app/infrastructure/db/sqlc"
	"github.com/kakkky/app/infrastructure/kvs"
)

func handleTask(mux *http.ServeMux) {
	sqlc := sqlc.NewSqlcQuerier()
	taskRepository := repository.NewTaskRepository(sqlc)
	taskQueryService := queryservice.NewTaskQueryService(sqlc)
	redisCom, err := kvs.NewRedisCommander()
	if err != nil {
		log.Fatalln(err)
	}
	authorization := middleware.Authorication(
		auth.NewAuthorizationUsecase(
			authInfra.NewJWTAuthenticator(),
			repository.NewTokenAuthenticatorRepository(redisCom),
		),
	)

	mux.Handle("POST /task", composeMiddlewares(authorization, middleware.Logger)(
		taskHandler.NewPostTaskHandler(
			taskUsecase.NewCreateTaskUsecase(
				taskRepository,
			),
		),
	))
	mux.Handle("DELETE /task/{id}", composeMiddlewares(authorization, middleware.Logger)(
		taskHandler.NewDeleteTaskHandler(
			taskUsecase.NewDeleteTaskUsecase(
				taskRepository,
			),
		),
	))
	mux.Handle("PATCH /task", composeMiddlewares(authorization, middleware.Logger)(
		taskHandler.NewUpdateTaskStateHandler(
			taskUsecase.NewUpdateTaskStateUsecase(
				taskRepository,
			),
		),
	))
	mux.Handle("GET /task/{id}", composeMiddlewares(authorization, middleware.Logger)(
		taskHandler.NewGetTaskHandler(
			taskUsecase.NewFetchTaskUsease(
				taskQueryService,
			),
		),
	))
	mux.Handle("GET /tasks", composeMiddlewares(authorization, middleware.Logger)(
		taskHandler.NewGetTasksHandler(
			taskUsecase.NewFetchTasksUsease(
				taskQueryService,
			),
		),
	))
	mux.Handle("GET /user/tasks", composeMiddlewares(authorization, middleware.Logger)(
		taskHandler.NewGetUserTasksHandler(
			taskUsecase.NewFetchUserTasksUsecase(
				taskQueryService,
			),
		),
	))
}
