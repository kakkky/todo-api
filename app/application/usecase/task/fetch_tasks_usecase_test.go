package task

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kakkky/app/domain/errors"
	"go.uber.org/mock/gomock"
)

func TestTask_FetchTasksUsecase_Run(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		mockFn  func(mq *MockTaskQueryService)
		errType error
		want    []*FetchTaskUsecaseOutputDTO
		wantErr bool
	}{
		{
			name: "正常系",
			mockFn: func(mq *MockTaskQueryService) {
				mq.EXPECT().FetchAllTasks(gomock.Any()).Return(
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
							UserName: "user2",
							UserId:   "user_id2",
							Content:  "content",
							State:    "todo",
						},
					}, nil,
				)
			},
			want: []*FetchTaskUsecaseOutputDTO{
				{
					ID:       "id",
					UserName: "user",
					UserId:   "user_id",
					Content:  "content",
					State:    "todo",
				},
				{
					ID:       "id2",
					UserName: "user2",
					UserId:   "user_id2",
					Content:  "content",
					State:    "todo",
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
			sut := NewFetchTasksUsease(mockTaskQueryService)
			ctx := context.Background()
			got, err := sut.Run(ctx)
			// 期待するエラー型を設定していた場合はエラー型を比較して検証する
			if tt.errType != nil && !errors.Is(err, tt.errType) {
				t.Errorf("fetchTasksUsecase.Run = error:%v,want errYType:%v", err, tt.errType)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchTasksUsecase.Run = error:%v,wantErr:%v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("fetchTasksUsecase.Run() -got,+want :%v ", diff)
			}
		})
	}
}
