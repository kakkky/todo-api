package user

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/app/domain/user"
	"go.uber.org/mock/gomock"
)

func TestUser_EditProfileUsecase_Run(t *testing.T) {
	tests := []struct {
		name    string
		mockFn  func(mr *user.MockUserRepository)
		input   EditProfileUsecaseInputDTO
		want    *EditProfileUsecaseOutputDTO
		wantErr bool
	}{
		{
			name: "正常形:プロフィールを編集できる",
			mockFn: func(mr *user.MockUserRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(user.ReconstructUser("1", "user@test.com", "testuser", ""), nil)
				mr.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

			},
			input: EditProfileUsecaseInputDTO{
				ID:    "1",
				Email: "updated@test.com",
				Name:  "updatedUser",
			},
			want: &EditProfileUsecaseOutputDTO{
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
			input: EditProfileUsecaseInputDTO{
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
			editProfileUsecase := NewEditProfileUsecase(mockUserRepository)
			ctx := context.Background()
			got, err := editProfileUsecase.Run(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("editProfileUsecase.Run = error:%v,wantErr:%v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("editProfileUsecase.Run() -got,+want :%v ", diff)
			}
		})
	}
}
