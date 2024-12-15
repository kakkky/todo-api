package user

import (
	"context"
	"testing"

	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/app/domain/user"
	"go.uber.org/mock/gomock"
)

func TestUser_UnregisterUsease_Run(t *testing.T) {
	tests := []struct {
		name    string
		mockFn  func(mr *user.MockUserRepository)
		input   UnregisterUsecaseInputDTO
		wantErr bool
	}{
		{
			name: "正常系: ユーザーを削除できる",
			mockFn: func(mr *user.MockUserRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(user.ReconstructUser("1", "", "", ""), nil)
				mr.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			},
			input: UnregisterUsecaseInputDTO{
				ID: "1",
			},
			wantErr: false,
		},
		{
			name: "準正常系: ユーザーが見つからない場合はエラーを返す",
			mockFn: func(mr *user.MockUserRepository) {
				mr.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(nil, errors.ErrNotFoundUser)
			},
			input: UnregisterUsecaseInputDTO{
				ID: "0",
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
			unregisterUsecase := NewUnregisterUsecase(mockUserRepository)
			ctx := context.Background()
			err := unregisterUsecase.Run(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("unregisterUsecase.Run = error:%v,wantErr:%v", err, tt.wantErr)
			}
		})
	}
}
