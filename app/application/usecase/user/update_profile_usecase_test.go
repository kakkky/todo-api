package user

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/app/domain/user"
	"go.uber.org/mock/gomock"
)

func TestUser_UpdateProfileUsecase_Run(t *testing.T) {
	tests := []struct {
		name    string
		mockFn  func(mr *user.MockUserRepository)
		input   UpdateProfileUsecaseInputDTO
		want    *UpdateProfileUsecaseOutputDTO
		wantErr bool
	}{
		{
			name: "正常形:プロフィールを編集できる",
			mockFn: func(mr *user.MockUserRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(user.ReconstructUser("1", "user@test.com", "testuser", ""), nil)
				mr.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

			},
			input: UpdateProfileUsecaseInputDTO{
				ID:    "1",
				Email: "updated@test.com",
				Name:  "updatedUser",
			},
			want: &UpdateProfileUsecaseOutputDTO{
				ID:    "1",
				Email: "updated@test.com",
				Name:  "updatedUser",
			},
			wantErr: false,
		},
		{
			name: "準正常形:存在しないユーザーは編集できない",
			mockFn: func(mr *user.MockUserRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(nil, errors.ErrNotFoundUser)
			},
			input: UpdateProfileUsecaseInputDTO{
				ID:    "0",
				Email: "noexistent@test.com",
				Name:  "noexitstent",
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
			mockUserRepository := user.NewMockUserRepository(ctrl)
			tt.mockFn(mockUserRepository)
			// ユースケースオブジェクト
			UpdateProfileUsecase := NewUpdateProfileUsecase(mockUserRepository)
			ctx := context.Background()
			got, err := UpdateProfileUsecase.Run(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateProfileUsecase.Run = error:%v,wantErr:%v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("UpdateProfileUsecase.Run() -got,+want :%v ", diff)
			}
		})
	}
}
