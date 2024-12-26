package task

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/app/domain/task"
	"go.uber.org/mock/gomock"
)

func TestTask_UpdateTaskStateUsecase_Run(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		mockFn  func(mr *task.MockTaskRepository)
		input   UpdateTaskStateUsecaseInputDTO
		errType error
		want    *UpdateTaskStateUsecaseOutputDTO
		wantErr bool
	}{
		{
			name: "正常系",
			mockFn: func(mr *task.MockTaskRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(
					task.ReconstructTask("id", "user_id", "this is content", 0),
					nil,
				)
				mr.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			},
			input: UpdateTaskStateUsecaseInputDTO{
				ID:    "id",
				State: "done",
			},
			want: &UpdateTaskStateUsecaseOutputDTO{
				ID:      "id",
				UserId:  "user_id",
				Content: "this is content",
				State:   "done",
			},
			wantErr: false,
		},
		{
			name: "準正常系:指定したstateが不正",
			mockFn: func(mr *task.MockTaskRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(
					task.ReconstructTask("id", "user_id", "this is content", 0),
					nil,
				)
			},
			errType: errors.ErrInvalidTaskState,
			input: UpdateTaskStateUsecaseInputDTO{
				ID:    "id",
				State: "invalid",
			},
			wantErr: true,
		},
		{
			name: "準正常系:指定したidを持つタスクが存在しない",
			mockFn: func(mr *task.MockTaskRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(
					nil,
					errors.ErrNotFoundTask,
				)
			},
			errType: errors.ErrNotFoundTask,
			input: UpdateTaskStateUsecaseInputDTO{
				ID:    "id",
				State: "doing",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// モックを設定
			ctrl := gomock.NewController(t)
			mockTaskRepository := task.NewMockTaskRepository(ctrl)
			tt.mockFn(mockTaskRepository)
			// ユースケースオブジェクト
			sut := NewUpdateTaskStateUsecase(mockTaskRepository)
			ctx := context.Background()
			got, err := sut.Run(ctx, tt.input)
			// 期待するエラー型を設定していた場合はエラー型を比較して検証する
			if tt.errType != nil && !errors.Is(err, tt.errType) {
				t.Errorf("updateTaskState.Run = error:%v,want errYType:%v", err, tt.errType)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("updateTaskState.Run = error:%v,wantErr:%v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("updateTaskState.Run() -got,+want :%v ", diff)
			}
		})
	}
}
