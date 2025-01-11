package queryservice

import "context"

// ORMはこのインターフェースを満たすように設計する
type Querier interface {
	FetchTaskById(ctx context.Context, id string) (FetchTaskByIdRow, error)
	FetchAllTasks(ctx context.Context) ([]FetchAllTasksRow, error)
	FetchUserTasks(ctx context.Context, userID string) ([]FetchUserTasksRow, error)
}

type FetchTaskByIdRow struct {
	ID      string
	Name    string
	UserID  string
	Content string
	State   int32
}

type FetchAllTasksRow struct {
	ID      string
	Name    string
	UserID  string
	Content string
	State   int32
}

type FetchUserTasksRow struct {
	ID      string
	Name    string
	UserID  string
	Content string
	State   int32
}
