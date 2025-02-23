package user

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kakkky/app/domain/user"
	"go.uber.org/mock/gomock"
)

func TestUser_FetchUserUsecase_Run(t *testing.T) {
	tests := []struct {
		name    string
		input   FetchUserUsecaseInputDTO
		mockFn  func(mr *user.MockUserRepository)
		want    *FetchUserUsecaseOutputDTO
		wantErr bool
	}{
		{
			name:  "正常形: idで指定したユーザーを取得する",
			input: FetchUserUsecaseInputDTO{ID: "1"},
			mockFn: func(mr *user.MockUserRepository) {
				mr.EXPECT().FindById(gomock.Any(), "1").Return(
					user.ReconstructUser("1", "", "user1", ""),
					nil,
				)
			},
			want: &FetchUserUsecaseOutputDTO{
				ID:   "1",
				Name: "user1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// モックを設定
			ctrl := gomock.NewController(t)
			mockUserRepository := user.NewMockUserRepository(ctrl)
			tt.mockFn(mockUserRepository)
			FetchUserUsecase := NewFetchUserUsecase(mockUserRepository)
			ctx := context.Background()
			got, err := FetchUserUsecase.Run(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchUserUsecase.Run = error:%v,wantErr:%v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FetchUserUsecase.Run() -got,+want :%v ", diff)
			}

		})
	}
}
