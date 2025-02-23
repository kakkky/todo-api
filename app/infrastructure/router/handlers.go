package router

import (
	"github.com/kakkky/app/adapter/presentation/handler/auth"
	"github.com/kakkky/app/adapter/presentation/handler/task"
	"github.com/kakkky/app/adapter/presentation/handler/user"
	"github.com/kakkky/app/adapter/queryservice"
	"github.com/kakkky/app/adapter/repository"
	authUsecase "github.com/kakkky/app/application/usecase/auth"
	taskUsecase "github.com/kakkky/app/application/usecase/task"
	userUsecase "github.com/kakkky/app/application/usecase/user"
	userDomain "github.com/kakkky/app/domain/user"
	authInfra "github.com/kakkky/app/infrastructure/auth"
	"github.com/kakkky/app/infrastructure/db/sqlc"
	"github.com/kakkky/app/infrastructure/kvs"
)

// 認証系ハンドラー
var (
	loginHandler  *auth.LoginHandler
	logoutHandler *auth.LogoutHandler
)

// ユーザー系ハンドラー
var (
	postUserHandler       *user.PostUserHandler
	deleteUserHandler     *user.DeleteUserHandler
	updateUserHandler     *user.UpdateUserHandler
	getUsersHandler       *user.GetUsersHandler
	getCurrentUserHandler *user.GetCurrentUserHandler
)

// タスク系ハンドラー
var (
	postTaskHandler        *task.PostTaskHandler
	deleteTaskHandler      *task.DeleteTaskHandler
	updateTaskStateHandler *task.UpdateTaskStateHandler
	getTaskHandler         *task.GetTaskHandler
	getTasksHandler        *task.GetTasksHandler
	getUserTasksHandler    *task.GetUserTasksHandler
)

// ハンドラーを初期化する
func initHandlers() {
	initAuthHandlers()
	initUserHandlers()
	initTaskHandlers()
}

func initAuthHandlers() {
	loginHandler = auth.NewLoginHandler(
		authUsecase.NewLoginUsecase(
			repository.NewUserRepository(sqlc.NewSqlcQuerier()),
			repository.NewJwtAuthenticatorRepository(kvs.NewRedisCommander()),
			authInfra.NewJwtAuthenticator(),
		),
	)

	logoutHandler = auth.NewLogoutHandler(
		authUsecase.NewLogoutUsecase(
			authInfra.NewJwtAuthenticator(),
			repository.NewJwtAuthenticatorRepository(kvs.NewRedisCommander()),
		),
	)
}

func initUserHandlers() {
	postUserHandler = user.NewPostUserHandler(
		userUsecase.NewRegisterUsecase(
			repository.NewUserRepository(sqlc.NewSqlcQuerier()),
			userDomain.NewUserDomainService(repository.NewUserRepository(sqlc.NewSqlcQuerier())),
		),
	)

	deleteUserHandler = user.NewDeleteUserHandler(
		userUsecase.NewUnregisterUsecase(
			repository.NewUserRepository(sqlc.NewSqlcQuerier()),
		),
	)

	getUsersHandler = user.NewGetUsersHandler(
		userUsecase.NewFetchUsersUsecase(
			repository.NewUserRepository(sqlc.NewSqlcQuerier()),
		),
	)

	getCurrentUserHandler = user.NewGetCurrentUserHandler(
		userUsecase.NewFetchUserUsecase(
			repository.NewUserRepository(sqlc.NewSqlcQuerier()),
		),
	)

	updateUserHandler = user.NewUpdateUserHandler(
		userUsecase.NewUpdateProfileUsecase(
			repository.NewUserRepository(sqlc.NewSqlcQuerier()),
		),
	)
}

func initTaskHandlers() {
	postTaskHandler = task.NewPostTaskHandler(
		taskUsecase.NewCreateTaskUsecase(
			repository.NewTaskRepository(sqlc.NewSqlcQuerier()),
		),
	)

	deleteTaskHandler = task.NewDeleteTaskHandler(
		taskUsecase.NewDeleteTaskUsecase(
			repository.NewTaskRepository(sqlc.NewSqlcQuerier()),
		),
	)

	updateTaskStateHandler = task.NewUpdateTaskStateHandler(
		taskUsecase.NewUpdateTaskStateUsecase(
			repository.NewTaskRepository(sqlc.NewSqlcQuerier()),
		),
	)

	getTaskHandler = task.NewGetTaskHandler(
		taskUsecase.NewFetchTaskUsease(
			queryservice.NewTaskQueryService(sqlc.NewSqlcQuerier()),
		),
	)

	getTasksHandler = task.NewGetTasksHandler(
		taskUsecase.NewFetchTasksUsease(
			queryservice.NewTaskQueryService(sqlc.NewSqlcQuerier()),
		),
	)

	getUserTasksHandler = task.NewGetUserTasksHandler(
		taskUsecase.NewFetchUserTasksUsecase(
			queryservice.NewTaskQueryService(sqlc.NewSqlcQuerier()),
		),
	)
}
