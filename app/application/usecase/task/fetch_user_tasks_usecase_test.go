package task

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kakkky/app/domain/errors"
	"go.uber.org/mock/gomock"
)

func TestTask_FetchUserTasksUsecase_Run(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		mockFn  func(mq *MockTaskQueryService)
		input   FetchUserTasksUsecaseInputDTO
		errType error
		want    []*FetchUserTasksUsecaseOutputDTO
		wantErr bool
	}{
		{
			name: "正常系",
			mockFn: func(mq *MockTaskQueryService) {
				mq.EXPECT().FetchTaskByUserId(gomock.Any(), gomock.Any()).Return(
					[]*FetchTaskDTO{
						{
							ID:       "id",
							UserName: "user",
							UserId:   "user_id",
							Content:  "content",
							State:    "todo",
						},
						{
							ID:       "id2",
							UserName: "user",
							UserId:   "user_id",
							Content:  "content",
							State:    "todo",
						},
					}, nil,
				)
			},
			input: FetchUserTasksUsecaseInputDTO{
				UserId: "user_id",
			},
			want: []*FetchUserTasksUsecaseOutputDTO{
				{
					ID:      "id",
					Content: "content",
					State:   "todo",
				},
				{
					ID:      "id2",
					Content: "content",
					State:   "todo",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockTaskQueryService := NewMockTaskQueryService(ctrl)
			tt.mockFn(mockTaskQueryService)
			// ユースケースオブジェクト
			sut := NewFetchUserTasksUsecase(mockTaskQueryService)
			ctx := context.Background()
			got, err := sut.Run(ctx, tt.input)
			// 期待するエラー型を設定していた場合はエラー型を比較して検証する
			if tt.errType != nil && !errors.Is(err, tt.errType) {
				t.Errorf("FetchUserTasksUsecase.Run = error:%v,want errYType:%v", err, tt.errType)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchUserTasksUsecase.Run = error:%v,wantErr:%v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FetchUserTasksUsecase.Run() -got,+want :%v ", diff)
			}
		})
	}
}
