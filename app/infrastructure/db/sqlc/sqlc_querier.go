package sqlc

import (
	"context"

	queryservice "github.com/kakkky/app/adapter/query_service"
	"github.com/kakkky/app/adapter/repository"
	"github.com/kakkky/app/infrastructure/db"
)

type SqlcQuerier struct {
	queries *Queries
}

func NewSqlcQuerier() *SqlcQuerier {
	return &SqlcQuerier{
		queries: New(db.GetDB()),
	}
}

func (sq *SqlcQuerier) DeleteTask(ctx context.Context, id string) error {
	return sq.queries.DeleteTask(ctx, id)
}

func (sq *SqlcQuerier) DeleteUser(ctx context.Context, id string) error {
	return sq.queries.DeleteUser(ctx, id)
}

func (sq *SqlcQuerier) FetchAllUser(ctx context.Context) ([]repository.FetchAllUsersRow, error) {
	rows, err := sq.queries.FetchAllUsers(ctx)
	var r []repository.FetchAllUsersRow
	for _, row := range rows {
		r = append(r, repository.FetchAllUsersRow{
			ID:             row.ID,
			Email:          row.Email,
			Name:           row.Name,
			HashedPassword: row.HashedPassword,
		})
	}
	return r, err
}

func (sq *SqlcQuerier) FindTaskById(ctx context.Context, id string) (repository.FindTaskByIdRow, error) {
	row, err := sq.queries.FindTaskById(ctx, id)
	return repository.FindTaskByIdRow(row), err
}

func (sq *SqlcQuerier) FindUserByEmail(ctx context.Context, email string) (repository.FindUserByEmailRow, error) {
	row, err := sq.queries.FindUserByEmail(ctx, email)
	return repository.FindUserByEmailRow(row), err
}

func (sq *SqlcQuerier) FindUserById(ctx context.Context, id string) (repository.FindUserByIdRow, error) {
	row, err := sq.queries.FindUserById(ctx, id)
	return repository.FindUserByIdRow(row), err
}

func (sq *SqlcQuerier) InsertTask(ctx context.Context, arg repository.InsertTaskParams) error {
	return sq.queries.InsertTask(ctx, InsertTaskParams(arg))
}

func (sq *SqlcQuerier) InsertUser(ctx context.Context, arg repository.InsertUserParams) error {
	return sq.queries.InsertUser(ctx, InsertUserParams(arg))
}

func (sq *SqlcQuerier) UpdateTask(ctx context.Context, arg repository.UpdateTaskParams) error {
	return sq.queries.UpdateTask(ctx, UpdateTaskParams(arg))
}

func (sq *SqlcQuerier) UpdateUser(ctx context.Context, arg repository.UpdateUserParams) error {
	return sq.queries.UpdateUser(ctx, UpdateUserParams(arg))
}

func (sq *SqlcQuerier) FetchTaskById(ctx context.Context, id string) (queryservice.FetchTaskByIdRow, error) {
	row, err := sq.queries.FetchTaskById(ctx, id)
	return queryservice.FetchTaskByIdRow(row), err
}

func (sq *SqlcQuerier) FetchAllTasks(ctx context.Context) ([]queryservice.FetchAllTasksRow, error) {
	rows, err := sq.queries.FetchAllTasks(ctx)
	var r []queryservice.FetchAllTasksRow
	for _, row := range rows {
		r = append(r, queryservice.FetchAllTasksRow(row))
	}
	return r, err
}

func (sq *SqlcQuerier) FetchUserTasks(ctx context.Context, userID string) ([]queryservice.FetchUserTasksRow, error) {
	rows, err := sq.queries.FetchUserTasks(ctx, userID)
	var r []queryservice.FetchUserTasksRow
	for _, row := range rows {
		r = append(r, queryservice.FetchUserTasksRow(row))
	}
	return r, err
}
