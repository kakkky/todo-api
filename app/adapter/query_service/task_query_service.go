package queryservice

import (
	"context"
	"database/sql"

	taskUsecase "github.com/kakkky/app/application/usecase/task"
	"github.com/kakkky/app/domain/errors"
	taskDomain "github.com/kakkky/app/domain/task"
	"github.com/kakkky/app/infrastructure/db/sqlc"
)

type taskQueryService struct{}

func NewTaskQueryService() taskUsecase.TaskQueryService {
	return &taskQueryService{}
}

func (tqs *taskQueryService) FetchTaskById(ctx context.Context, id string) (*taskUsecase.FetchTaskDTO, error) {
	queries := sqlc.GetQueries()
	r, err := queries.FetchTaskById(ctx, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, errors.ErrNotFoundTask
	}
	if err != nil {
		return nil, err
	}
	// Stateをint→strにする必要があり、
	// State型のメソッドを使用するために再構成する
	t := taskDomain.ReconstructTask(
		r.ID,
		r.UserID,
		r.Content,
		int(r.State),
	)
	return &taskUsecase.FetchTaskDTO{
		ID:       t.GetID(),
		UserName: r.Name,
		UserId:   t.GetUserId(),
		Content:  t.GetContent().Value(),
		State:    t.GetState().StrValue(),
	}, nil

}

func (tqs *taskQueryService) FetchUserTasks(ctx context.Context, userId string) ([]*taskUsecase.FetchTaskDTO, error) {
	queries := sqlc.GetQueries()
	rs, err := queries.FetchUserTasks(ctx, userId)
	if err != nil {
		return nil, err
	}
	dtos := make([]*taskUsecase.FetchTaskDTO, 0, len(rs))
	for _, r := range rs {
		t := taskDomain.ReconstructTask(
			r.ID,
			r.UserID,
			r.Content,
			int(r.State),
		)
		dtos = append(dtos, &taskUsecase.FetchTaskDTO{
			ID:       t.GetID(),
			UserName: r.Name,
			UserId:   t.GetUserId(),
			Content:  t.GetContent().Value(),
			State:    t.GetState().StrValue(),
		})
	}
	return dtos, nil
}

func (tqs *taskQueryService) FetchTasks(ctx context.Context) ([]*taskUsecase.FetchTaskDTO, error) {
	queries := sqlc.GetQueries()
	rs, err := queries.FetchTasks(ctx)
	if err != nil {
		return nil, err
	}
	dtos := make([]*taskUsecase.FetchTaskDTO, 0, len(rs))
	for _, r := range rs {
		t := taskDomain.ReconstructTask(
			r.ID,
			r.UserID,
			r.Content,
			int(r.State),
		)
		dtos = append(dtos, &taskUsecase.FetchTaskDTO{
			ID:       t.GetID(),
			UserName: r.Name,
			UserId:   t.GetUserId(),
			Content:  t.GetContent().Value(),
			State:    t.GetState().StrValue(),
		})
	}
	return dtos, nil
}
