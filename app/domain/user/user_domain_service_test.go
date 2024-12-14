package user

import (
	"context"
	"testing"

	"github.com/kakkky/app/domain/errors"
	gomock "go.uber.org/mock/gomock"
)

func TestUserDomainService_IsExists(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		email   Email
		mockFn  func(m *MockUserRepository)
		want    bool
		wantErr bool
	}{
		{
			name:  "正常系: ユーザー存在する",
			email: Email{value: "test@example.com"},
			mockFn: func(m *MockUserRepository) {
				// ErrNotFoundUserを返す
				m.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(&User{}, nil)
			},
			want:    true,
			wantErr: false,
		},
		{
			name:  "準正常系: ユーザーが存在しない",
			email: Email{value: "test@example.com"},
			mockFn: func(m *MockUserRepository) {
				m.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(nil, errors.ErrNotFoundUser)
			},
			want:    false,
			wantErr: false,
		},
		{
			name:  "異常系:リポジトリからErrNotFoundUser以外のエラーが返る場合",
			email: Email{value: "test@example.com"},
			mockFn: func(m *MockUserRepository) {
				m.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(nil, errors.New("Database error"))
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mockUserRepository := NewMockUserRepository(ctrl)
			// モック呼び出しの設定
			tt.mockFn(mockUserRepository)
			// ドメインサービスを作成
			sut := NewUserDomainService(mockUserRepository)
			ctx := context.Background()
			got, err := sut.IsExists(ctx, tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("userDomainService.IsExists() want error %v,but got %v : %v", tt.wantErr, got, err)
			}
			if got != tt.want {
				t.Errorf("userDomainService.IsExists()= %v,but want %v", got, tt.want)
			}

		})
	}
}
