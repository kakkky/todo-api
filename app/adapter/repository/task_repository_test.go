package repository

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/app/domain/task"
	testhelper "github.com/kakkky/app/infrastructure/db/testhelper"
)

func TestTaskRepository_Save(t *testing.T) {
	taskRepository := NewTaskRepository()
	ta, _ := task.NewTask(
		"1",
		"content",
		"todo",
	)
	type args struct {
		task *task.Task
	}
	tests := []struct {
		name    string
		args    args
		want    *task.Task
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				task: ta,
			},
			want:    ta,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			testhelper.SetupFixtures("testdata/fixtures/users.yml")

			ctx := context.Background()
			// タスクを保存
			if err := taskRepository.Save(ctx, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("taskRepository.Save() = error %v, wantErr %v", err, tt.wantErr)
			}
			// FindByIdで検索して確かめる
			got, err := taskRepository.FindById(ctx, ta.GetID())
			if (err != nil) != tt.wantErr {
				t.Errorf("taskRepository.FindById() =　error %v,but wantErr :%v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(task.Task{}, task.Content{}, task.State{}), cmpopts.IgnoreFields(task.Task{}, "id")); diff != "" {
				t.Errorf("userRepository.FindByEmail() -got,+want :%v ", diff)
			}
		})
	}
}

func TestTaskRepository_FindById(t *testing.T) {
	taskRepository := NewTaskRepository()
	task1 := task.ReconstructTask(
		"1",
		"1",
		"content",
		0,
	)

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *task.Task
		errType error
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				id: "1",
			},
			want:    task1,
			wantErr: false,
		},
		{
			name: "準正常系: 存在しないタスクIDではErrNotFoundTaskが返る",
			args: args{
				id: "invalid",
			},
			errType: errors.ErrNotFoundTask,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// task1のみがDBに保存されている
			testhelper.SetupFixtures("testdata/fixtures/users.yml", "testdata/fixtures/tasks.yml")
			ctx := context.Background()
			got, err := taskRepository.FindById(ctx, tt.args.id)
			if (err != nil) != tt.wantErr && tt.errType != nil && errors.Is(err, tt.errType) {
				t.Errorf("taskRepository.FindById() = error %v, want errType %v", err, tt.errType)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("taskRepository.FindById() = error %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(task.Task{}, task.Content{}, task.State{})); diff != "" {
				t.Errorf("taskRepository.FindById() -got,+want : %v", diff)
			}
		})
	}
}

func TestTaskRepository_Update(t *testing.T) {
	taskRepository := NewTaskRepository()
	// 更新するタスク
	updatingTask := task.ReconstructTask(
		"1",
		"1",
		"content",
		1,
	)

	type args struct {
		task *task.Task
	}
	tests := []struct {
		name    string
		args    args
		want    *task.Task
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				task: updatingTask,
			},
			want:    updatingTask,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// 既存のタスクを更新
			testhelper.SetupFixtures("testdata/fixtures/users.yml", "testdata/fixtures/tasks.yml")

			ctx := context.Background()
			if err := taskRepository.Update(ctx, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("taskRepository.Update() = error %v, wantErr %v", err, tt.wantErr)
			}
			// FindByIdで検索して確かめる
			got, err := taskRepository.FindById(ctx, tt.args.task.GetID())
			if (err != nil) != tt.wantErr {
				t.Errorf("taskRepository.FindById() = error %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(task.Task{}, task.Content{}, task.State{})); diff != "" {
				t.Errorf("taskRepository.FindById() -got,+want : %v", diff)
			}
		})
	}
}

func TestTaskRepository_Delete(t *testing.T) {
	taskRepository := NewTaskRepository()
	// 更新するタスク
	deletingTask := task.ReconstructTask(
		"1",
		"1",
		"content",
		1,
	)
	type args struct {
		task *task.Task
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				deletingTask,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// 既存のタスクを更新
			testhelper.SetupFixtures("testdata/fixtures/users.yml", "testdata/fixtures/tasks.yml")

			ctx := context.Background()
			if err := taskRepository.Delete(ctx, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("taskRepository.Delete() = error %v, wantErr %v", err, tt.wantErr)
			}
			// FindByIdで検索して確かめる
			_, err := taskRepository.FindById(ctx, tt.args.task.GetID())
			if (err != nil) != tt.wantErr && !errors.Is(err, errors.ErrNotFoundTask) {
				t.Errorf("taskRepository.FindById()=error:%v, wantErr:%v", err, tt.wantErr)
			}
		})
	}
}
