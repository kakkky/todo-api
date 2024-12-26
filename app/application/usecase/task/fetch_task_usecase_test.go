package task

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kakkky/app/domain/errors"
	"go.uber.org/mock/gomock"
)

func TestTask_FetchTaskUsecase_Run(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		mockFn  func(mq *MockTaskQueryService)
		input   FetchTaskUsecaseInputDTO
		errType error
		want    *FetchTaskUsecaseOutputDTO
		wantErr bool
	}{
		{
			name: "正常系",
			mockFn: func(mq *MockTaskQueryService) {
				mq.EXPECT().FetchTaskById(gomock.Any(), gomock.Any()).Return(
					&FetchTaskDTO{
						ID:       "id",
						UserName: "user",
						UserId:   "user_id",
						Content:  "content",
						State:    "todo",
					}, nil,
				)
			},
			input: FetchTaskUsecaseInputDTO{
				ID: "id",
			},
			want: &FetchTaskUsecaseOutputDTO{
				ID:       "id",
				UserName: "user",
				UserId:   "user_id",
				Content:  "content",
				State:    "todo",
			},
			wantErr: false,
		},
		{
			name: "準正常系：指定のidを持つtaskが見つからない",
			mockFn: func(mq *MockTaskQueryService) {
				mq.EXPECT().FetchTaskById(gomock.Any(), gomock.Any()).Return(
					nil, errors.ErrNotFoundTask,
				)
			},
			input: FetchTaskUsecaseInputDTO{
				ID: "id",
			},
			errType: errors.ErrNotFoundTask,
			wantErr: true,
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
			sut := NewFetchTaskUsease(mockTaskQueryService)
			ctx := context.Background()
			got, err := sut.Run(ctx, tt.input)
			// 期待するエラー型を設定していた場合はエラー型を比較して検証する
			if tt.errType != nil && !errors.Is(err, tt.errType) {
				t.Errorf("fetchTaskUsecase.Run = error:%v,want errYType:%v", err, tt.errType)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchTaskUsecase.Run = error:%v,wantErr:%v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("fetchTaskUsecase.Run() -got,+want :%v ", diff)
			}
		})
	}
}
