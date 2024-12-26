package task

import (
	"context"

	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/app/domain/task"
	"go.uber.org/mock/gomock"
)

func TestTask_CreateTaskUsecase_Run(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		mockFn  func(mr *task.MockTaskRepository)
		input   CreateTaskUsecaseInputDTO
		errType error
		want    *CreateTaskUsecaseOutputDTO
		wantErr bool
	}{
		{
			name: "正常系",
			mockFn: func(mr *task.MockTaskRepository) {
				mr.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
			},
			input: CreateTaskUsecaseInputDTO{
				UserId:  "user_id",
				Content: "this is content",
				State:   "todo",
			},
			want: &CreateTaskUsecaseOutputDTO{
				UserId:  "user_id",
				Content: "this is content",
				State:   "todo",
			},
			wantErr: false,
		},
		{
			name: "準正常系：タスク内容が空っぽ",
			mockFn: func(mr *task.MockTaskRepository) {
				// Saveメソッドは呼ばれない
			},
			input: CreateTaskUsecaseInputDTO{
				UserId:  "user_id",
				Content: "",
				State:   "todo",
			},
			errType: errors.ErrContentEmpty,
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
			sut := NewCreateTaskUsecase(mockTaskRepository)
			ctx := context.Background()
			got, err := sut.Run(ctx, tt.input)
			// 期待するエラー型を設定していた場合はエラー型を比較して検証する
			if tt.errType != nil && !errors.Is(err, tt.errType) {
				t.Errorf("createTaskUsecase.Run = error:%v,want errYType:%v", err, tt.errType)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("createTaskUsecase.Run = error:%v,wantErr:%v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(CreateTaskUsecaseOutputDTO{}, "ID")); diff != "" {
				t.Errorf("createTaskUsecase.Run() -got,+want :%v ", diff)
			}
		})
	}
}
