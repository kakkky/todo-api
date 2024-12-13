package repository

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kakkky/app/domain/user"
)

func TestUserRepository_Save_And_FindByEmail(t *testing.T) {
	t.Parallel()
	// ユーザーインスタンスを用意
	user1, _ := user.NewUser(
		"user1@example.com",
		"user1",
		"password",
	)
	userRepository := NewUserRepository()
	type arg struct {
		user  *user.User
		email user.Email
	}
	// テストテーブル
	tests := []struct {
		name    string
		arg     arg
		want    *user.User
		wantErr bool
	}{
		{
			name: "正常系: 1人ユーザーが追加され、emailで検索できる",
			arg: arg{
				user:  user1,
				email: user1.GetEmail(),
			},
			want:    user1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			// ユーザーを保存
			if err := userRepository.Save(ctx, tt.arg.user); (err != nil) != tt.wantErr {
				t.Errorf("userRepository.Save()=error %v ,but wantErr:%v", err, tt.wantErr)
			}
			// ユーザーをemailで検索
			got, err := userRepository.FindByEmail(ctx, tt.arg.email)
			// 返すエラーはErrNotFoundUserのみであるべき
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FindByEmail() =　error %v,but wantErr :%v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(user.User{}, user.Email{}, user.HashedPassword{}), cmpopts.IgnoreFields(user.User{}, "id", "hashedPassword")); diff != "" {
				t.Errorf("userRepository.FindByEmail() -got,+want :%v ", diff)
			}
		})
	}
}
