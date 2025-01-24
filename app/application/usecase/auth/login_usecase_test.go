package auth

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kakkky/app/domain/user"
	"go.uber.org/mock/gomock"
)

func TestAuth_LoginUsecase_Run(t *testing.T) {
	t.Parallel()
	user1, _ := user.NewUser("user1@test.com", "user1", "password")
	tests := []struct {
		name    string
		mockFn  func(mur *user.MockUserRepository, ma *MockJwtAuthenticator, mar *MockJwtAuthenticatorRepository)
		input   LoginUsecaseInputDTO
		want    *LoginUsecaseOutputDTO
		wantErr bool
	}{
		{
			name: "正常系:ユーザーにJWTトークンが返る",
			mockFn: func(mur *user.MockUserRepository, ma *MockJwtAuthenticator, mar *MockJwtAuthenticatorRepository) {
				mur.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(user1, nil)
				ma.EXPECT().GenerateJwtToken(gomock.Any(), gomock.Any()).Return("signedToken", nil)
				mar.EXPECT().Save(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			input: LoginUsecaseInputDTO{
				Email:    "user1@test.com",
				Password: "password",
			},
			want: &LoginUsecaseOutputDTO{
				SignedToken: "signedToken",
			},
			wantErr: false,
		},
		{
			name: "準正常系:パスワードが異なる",
			mockFn: func(mur *user.MockUserRepository, ma *MockJwtAuthenticator, mar *MockJwtAuthenticatorRepository) {
				mur.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(user1, nil)
			},
			input: LoginUsecaseInputDTO{
				Email:    "user1@test.com",
				Password: "invalid",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUserRepository := user.NewMockUserRepository(ctrl)
			mockAuthenticator := NewMockJwtAuthenticator(ctrl)
			mockAuthenticatorRepository := NewMockJwtAuthenticatorRepository(ctrl)
			tt.mockFn(mockUserRepository, mockAuthenticator, mockAuthenticatorRepository)
			sut := NewLoginUsecase(mockUserRepository, mockAuthenticatorRepository, mockAuthenticator)
			ctx := context.Background()
			got, err := sut.Run(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginUsecase.Run error=%v,but wantErr%v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("LoginUsecase.Run -got,+want :%v", diff)
			}
		})
	}
}
