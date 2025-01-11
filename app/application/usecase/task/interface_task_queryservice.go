package task

import "context"

//go:generate mockgen -package=task -source=./interface_task_queryservice.go -destination=./mock_task_queryservice.go
type TaskQueryService interface {
	FetchTaskById(ctx context.Context, id string) (*FetchTaskDTO, error)
	FetchUserTasks(ctx context.Context, userId string) ([]*FetchTaskDTO, error)
	FetchAllTasks(ctx context.Context) ([]*FetchTaskDTO, error)
}

type FetchTaskDTO struct {
	ID       string
	UserName string
	UserId   string
	Content  string
	State    string
}
