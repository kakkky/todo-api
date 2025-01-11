package repository

import "context"

// ORMはこのインターフェースを満たすように設計する
type Querier interface {
	DeleteTask(ctx context.Context, id string) error
	DeleteUser(ctx context.Context, id string) error
	FetchAllUser(ctx context.Context) ([]FetchAllUsersRow, error)
	FindTaskById(ctx context.Context, id string) (FindTaskByIdRow, error)
	FindUserByEmail(ctx context.Context, email string) (FindUserByEmailRow, error)
	FindUserById(ctx context.Context, id string) (FindUserByIdRow, error)
	InsertTask(ctx context.Context, arg InsertTaskParams) error
	InsertUser(ctx context.Context, arg InsertUserParams) error
	UpdateTask(ctx context.Context, arg UpdateTaskParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

type FetchAllUsersRow struct {
	ID             string
	Email          string
	Name           string
	HashedPassword string
}

type FindUserByEmailRow struct {
	ID             string
	Email          string
	Name           string
	HashedPassword string
}

type FindUserByIdRow struct {
	ID             string
	Email          string
	Name           string
	HashedPassword string
}

type InsertUserParams struct {
	ID             string
	Name           string
	Email          string
	HashedPassword string
}

type UpdateUserParams struct {
	Name  string
	Email string
	ID    string
}

type FindTaskByIdRow struct {
	ID      string
	UserID  string
	Content string
	State   int32
}

type InsertTaskParams struct {
	ID      string
	UserID  string
	Content string
	State   int32
}

type UpdateTaskParams struct {
	State int32
	ID    string
}
