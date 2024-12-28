package queryservice

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kakkky/app/application/usecase/task"
	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/app/infrastructure/db/testhelper"
)

func TestTaskQueryService_FetchTaskById(t *testing.T) {
	taskQueryService := NewTaskQueryService()
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *task.FetchTaskDTO
		errType error
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				id: "1",
			},
			want: &task.FetchTaskDTO{
				ID:       "1",
				UserName: "user1",
				UserId:   "1",
				Content:  "content",
				State:    "todo",
			},
			wantErr: false,
		},
		{
			name: "準正常系：指定idのタスクが見つからない",
			args: args{
				id: "invalid",
			},
			errType: errors.ErrNotFoundTask,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testhelper.SetupFixtures("testdata/fixtures/users.yml", "testdata/fixtures/tasks.yml")
			ctx := context.Background()
			got, err := taskQueryService.FetchTaskById(ctx, tt.args.id)
			if (err != nil) != tt.wantErr && tt.errType != nil && errors.Is(err, tt.errType) {
				t.Errorf("taskQueryService.FetchTaskById = error %v, want errType %v", err, tt.errType)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("taskQueryService.FetchTaskById = error %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("taskQueryService.FetchTaskById() -got,+want : %v", diff)
			}
		})
	}
}

func TestTaskQueryService_FetchUserTasks(t *testing.T) {
	taskQueryService := NewTaskQueryService()
	type args struct {
		user_id string
	}
	tests := []struct {
		name    string
		args    args
		want    []*task.FetchTaskDTO
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				user_id: "1",
			},
			want: []*task.FetchTaskDTO{
				{
					ID:       "1",
					UserName: "user1",
					UserId:   "1",
					Content:  "content",
					State:    "todo",
				},
				{
					ID:       "2",
					UserName: "user1",
					UserId:   "1",
					Content:  "content",
					State:    "todo",
				},
				{
					ID:       "3",
					UserName: "user1",
					UserId:   "1",
					Content:  "content",
					State:    "doing",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testhelper.SetupFixtures("testdata/fixtures/users.yml", "testdata/fixtures/tasks.yml")
			ctx := context.Background()
			got, err := taskQueryService.FetchUserTasks(ctx, tt.args.user_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("taskQueryService.FetchUserTasks = error %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("taskQueryService.FetchUserTasks) -got,+want : %v", diff)
			}
		})
	}
}

func TestTaskQueryService_FetchTasks(t *testing.T) {
	taskQueryService := NewTaskQueryService()
	tests := []struct {
		name    string
		want    []*task.FetchTaskDTO
		wantErr bool
	}{
		{
			name: "正常系",
			want: []*task.FetchTaskDTO{
				{
					ID:       "1",
					UserName: "user1",
					UserId:   "1",
					Content:  "content",
					State:    "todo",
				},
				{
					ID:       "2",
					UserName: "user1",
					UserId:   "1",
					Content:  "content",
					State:    "todo",
				},
				{
					ID:       "3",
					UserName: "user1",
					UserId:   "1",
					Content:  "content",
					State:    "doing",
				},
				{
					ID:       "4",
					UserName: "user2",
					UserId:   "2",
					Content:  "content",
					State:    "doing",
				},
				{
					ID:       "5",
					UserName: "user2",
					UserId:   "2",
					Content:  "content",
					State:    "done",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testhelper.SetupFixtures("testdata/fixtures/users.yml", "testdata/fixtures/tasks.yml")
			ctx := context.Background()
			got, err := taskQueryService.FetchTasks(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("taskQueryService.FetchUserTasks = error %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("taskQueryService.FetchUserTasks) -got,+want : %v", diff)
			}
		})
	}
}
