package task

import "context"

//go:generate mockgen -package=task -source=./interface_task_query_service.go -destination=./mock_task_query_service.go
type TaskQueryService interface {
	FetchTaskById(ctx context.Context, id string) (*FetchTaskDTO, error)
	FetchTasks(ctx context.Context) ([]*FetchTaskDTO, error)
}

type FetchTaskDTO struct {
	ID      string
	Name    string
	Content string
	State   string
}
