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
		mockFn  func(mr *user.MockUserRepository)
		want    []*FetchUsersUsecaseOutputDTO
		wantErr bool
	}{
		{
			name: "正常形: 登録されている全てのユーザーを取得する",
			mockFn: func(mr *user.MockUserRepository) {
				mr.EXPECT().FetchAllUsers(gomock.Any()).Return(
					user.Users{
						user.ReconstructUser("1", "", "user1", ""),
						user.ReconstructUser("2", "", "user2", ""),
						user.ReconstructUser("3", "", "user3", ""),
					}, nil,
				)
			},
			want: []*FetchUsersUsecaseOutputDTO{
				{
					ID:   "1",
					Name: "user1",
				},
				{
					ID:   "2",
					Name: "user2",
				},
				{
					ID:   "3",
					Name: "user3",
				},
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
			FetchUsersUsecase := NewFetchUsersUsecase(mockUserRepository)
			ctx := context.Background()
			got, err := FetchUsersUsecase.Run(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchUsersUsecase.Run = error:%v,wantErr:%v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FetchUsersUsecase.Run() -got,+want :%v ", diff)
			}

		})
	}
}
