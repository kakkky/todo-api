package router

import (
	"github.com/kakkky/app/adapter/presentation/handler/auth"
	"github.com/kakkky/app/adapter/presentation/handler/task"
	"github.com/kakkky/app/adapter/presentation/handler/user"
	authUsecase "github.com/kakkky/app/application/usecase/auth"
	taskUsecase "github.com/kakkky/app/application/usecase/task"
	userUsecase "github.com/kakkky/app/application/usecase/user"
	userDomain "github.com/kakkky/app/domain/user"
)

// 認証系ハンドラー
var (
	loginHandler  *auth.LoginHandler
	logoutHandler *auth.LogoutHandler
)

// ユーザー系ハンドラー
var (
	postUserHandler   *user.PostUserHandler
	deleteUserHandler *user.DeleteUserHandler
	updateUserHandler *user.UpdateUserHandler
	getUsersHandler   *user.GetUsersHandler
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
			userRepository,
			jwtAuthenticatorRepository,
			jwtAuthenticator,
		),
	)

	logoutHandler = auth.NewLogoutHandler(
		authUsecase.NewLogoutUsecase(
			jwtAuthenticator,
			jwtAuthenticatorRepository,
		),
	)
}

func initUserHandlers() {
	postUserHandler = user.NewPostUserHandler(
		userUsecase.NewRegisterUsecase(
			userRepository,
			userDomain.NewUserDomainService(userRepository),
		),
	)

	deleteUserHandler = user.NewDeleteUserHandler(
		userUsecase.NewUnregisterUsecase(
			userRepository,
		),
	)

	getUsersHandler = user.NewGetUsersHandler(
		userUsecase.NewFetchUsersUsecase(
			userRepository,
		),
	)

	updateUserHandler = user.NewUpdateUserHandler(
		userUsecase.NewUpdateProfileUsecase(
			userRepository,
		),
	)
}

func initTaskHandlers() {
	postTaskHandler = task.NewPostTaskHandler(
		taskUsecase.NewCreateTaskUsecase(
			taskRepository,
		),
	)

	deleteTaskHandler = task.NewDeleteTaskHandler(
		taskUsecase.NewDeleteTaskUsecase(
			taskRepository,
		),
	)

	updateTaskStateHandler = task.NewUpdateTaskStateHandler(
		taskUsecase.NewUpdateTaskStateUsecase(
			taskRepository,
		),
	)

	getTaskHandler = task.NewGetTaskHandler(
		taskUsecase.NewFetchTaskUsease(
			taskQueryService,
		),
	)

	getTasksHandler = task.NewGetTasksHandler(
		taskUsecase.NewFetchTasksUsease(
			taskQueryService,
		),
	)

	getUserTasksHandler = task.NewGetUserTasksHandler(
		taskUsecase.NewFetchUserTasksUsecase(
			taskQueryService,
		),
	)
}
