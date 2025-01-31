package router

import (
	"github.com/kakkky/app/adapter/queryservice"
	"github.com/kakkky/app/adapter/repository"
	"github.com/kakkky/app/application/usecase/auth"
	"github.com/kakkky/app/application/usecase/task"
	taskDomain "github.com/kakkky/app/domain/task"
	"github.com/kakkky/app/domain/user"
	authInfra "github.com/kakkky/app/infrastructure/auth"
	"github.com/kakkky/app/infrastructure/db/sqlc"
	"github.com/kakkky/app/infrastructure/kvs"
)

// ハンドラーの初期化に必要なコンポーネント群
var sqlcQuerier *sqlc.SqlcQuerier
var userRepository user.UserRepository
var (
	taskRepository   taskDomain.TaskRepository
	taskQueryService task.TaskQueryService
)
var (
	jwtAuthenticatorRepository auth.JwtAuthenticatorRepository
	jwtAuthenticator           auth.JwtAuthenticator
)

// ハンドラーの初期化に必要なコンポーネント群を初期化する
func initDependencies() {
	sqlcQuerier = sqlc.NewSqlcQuerier()
	userRepository = repository.NewUserRepository(sqlcQuerier)
	taskRepository = repository.NewTaskRepository(sqlcQuerier)
	taskQueryService = queryservice.NewTaskQueryService(sqlcQuerier)
	jwtAuthenticatorRepository = repository.NewJwtAuthenticatorRepository(kvs.NewRedisCommander())
	jwtAuthenticator = authInfra.NewJwtAuthenticator()
}
