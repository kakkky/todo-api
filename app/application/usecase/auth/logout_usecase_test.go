package auth

import (
	"context"
	"fmt"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestLogoutUsecase(t *testing.T) {
	tests := []struct {
		name    string
		mockFn  func(ma *MockJwtAuthenticator, mar *MockJwtAuthenticatorRepository)
		input   LogoutUsecaseInputDTO
		wantErr bool
	}{
		{
			name: "正常系:redisにあるレコードを削除してログアウトする",
			mockFn: func(ma *MockJwtAuthenticator, mar *MockJwtAuthenticatorRepository) {
				mar.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			},
			input: LogoutUsecaseInputDTO{
				UserID: "1",
			},
			wantErr: false,
		},
		{
			name: "異常系:redis内でエラーが起きた場合、エラーを返す",
			mockFn: func(ma *MockJwtAuthenticator, mar *MockJwtAuthenticatorRepository) {
				mar.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(fmt.Errorf("err"))
			},
			input: LogoutUsecaseInputDTO{
				UserID: "1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockAuthenticator := NewMockJwtAuthenticator(ctrl)
			mockAuthenticatorRepository := NewMockJwtAuthenticatorRepository(ctrl)
			tt.mockFn(mockAuthenticator, mockAuthenticatorRepository)
			sut := NewLogoutUsecase(mockAuthenticator, mockAuthenticatorRepository)
			ctx := context.Background()
			err := sut.Run(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogoutUsecase.Run error=%v,but wantErr%v", err, tt.wantErr)
			}
		})
	}
}
