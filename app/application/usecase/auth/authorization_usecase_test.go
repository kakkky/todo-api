package auth

import (
	"context"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
)

func TestAuthorizationUsecase_Run(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		input   AuthorizationInputDTO
		mockFn  func(ma *MockJwtAuthenticator, mar *MockJwtAuthenticatorRepository)
		want    *AuthorizationOutputDTO
		wantErr bool
	}{
		{
			name: "正常系：トークンの認証に成功し、userIDを返す",
			input: AuthorizationInputDTO{
				SignedToken: "signedToken",
			},
			mockFn: func(ma *MockJwtAuthenticator, mar *MockJwtAuthenticatorRepository) {
				ma.EXPECT().VerifyToken(gomock.Any()).Return(&jwt.Token{}, nil)
				ma.EXPECT().VerifyExpiresAt(&jwt.Token{}).Return(nil)
				ma.EXPECT().GetJWTIDFromClaim(&jwt.Token{}).Return("jti", nil)
				ma.EXPECT().GetSubFromClaim(&jwt.Token{}).Return("userID", nil)
				mar.EXPECT().Load(gomock.Any(), "userID").Return("jti", nil)
			},
			want: &AuthorizationOutputDTO{
				UserID: "userID",
			},
			wantErr: false,
		},
		{
			name: "準正常系：kvsから得たjtiとトークンから得たjtiが一致しない",
			input: AuthorizationInputDTO{
				SignedToken: "signedToken",
			},
			mockFn: func(ma *MockJwtAuthenticator, mar *MockJwtAuthenticatorRepository) {
				ma.EXPECT().VerifyToken(gomock.Any()).Return(&jwt.Token{}, nil)
				ma.EXPECT().VerifyExpiresAt(&jwt.Token{}).Return(nil)
				ma.EXPECT().GetJWTIDFromClaim(&jwt.Token{}).Return("jti", nil)
				ma.EXPECT().GetSubFromClaim(&jwt.Token{}).Return("userID", nil)
				mar.EXPECT().Load(gomock.Any(), "userID").Return("", nil)
			},
			want: &AuthorizationOutputDTO{
				UserID: "userID",
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

			sut := NewAuthorizationUsecase(mockAuthenticator, mockAuthenticatorRepository)
			ctx := context.Background()
			got, err := sut.Run(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthorizationUsecase.Run error=%v,but wantErr%v", err, tt.wantErr)
			}
			// エラーが起きた場合はoutputを返さない
			if got == nil {
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AuthorizationUsecase.Run -got,+want :%v", diff)
			}
		})
	}
}
