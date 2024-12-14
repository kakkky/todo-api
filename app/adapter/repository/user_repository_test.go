package repository

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/app/domain/user"
	testhelper "github.com/kakkky/app/infrastructure/db/test_helper"
)

func TestUserRepository_Save_And_FindByEmail(t *testing.T) {
	t.Parallel()
	userRepository := NewUserRepository()
	// ユーザーインスタンスを用意
	user0, _ := user.NewUser(
		"user0@example.com",
		"user0",
		"password",
	)
	type args struct {
		user  *user.User
		email user.Email
	}
	// テストテーブル
	tests := []struct {
		name    string
		args    args
		want    *user.User
		wantErr bool
	}{
		{
			name: "正常系: 1人ユーザーが追加され、emailで検索できる",
			args: args{
				user:  user0,
				email: user0.GetEmail(),
			},
			want:    user0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			// ユーザーを保存
			if err := userRepository.Save(ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userRepository.Save()=error %v ,but wantErr:%v", err, tt.wantErr)
			}
			// ユーザーをemailで検索
			got, err := userRepository.FindByEmail(ctx, tt.args.email)
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

func TestUserRepository_FindByEmail(t *testing.T) {
	t.Parallel()
	userRepository := NewUserRepository()
	user1, _ := user.NewUser(
		"user1@test.com",
		"user1",
		"password",
	)
	user2, _ := user.NewUser(
		"user2@test.com",
		"user2",
		"password",
	)
	type args struct {
		email user.Email
	}
	tests := []struct {
		name    string
		args    args
		want    *user.User
		errType error
		wantErr bool
	}{
		{
			name: "正常系：ユーザーを検索できる",
			args: args{
				email: user1.GetEmail(),
			},
			want:    user1,
			wantErr: false,
		},
		{
			name: "準正常系：ユーザーが見つからなければErrNotFoundUserが返ってくる",
			args: args{
				// 存在しないuser2のemailで検索
				email: user2.GetEmail(),
			},
			errType: errors.ErrNotFoundUser,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// user1のみがDBに保存されている
			testhelper.SetupFixtures(t, "testdata/fixtures/users/users.yml")
			ctx := context.Background()
			got, err := userRepository.FindByEmail(ctx, tt.args.email)
			// 期待されるエラータイプが設定されている場合はそれも検証
			if (err != nil) != tt.wantErr && tt.errType != nil {
				t.Errorf("userRepository.FindByEmail() =error:%v, want errType:%v", err, tt.errType)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FindByEmail() =error:%v, wantErr:%v", err, err)
			}
			if diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(user.User{}, user.Email{}, user.HashedPassword{}), cmpopts.IgnoreFields(user.User{}, "id", "hashedPassword")); diff != "" {
				t.Errorf("userRepository.FindByEmail() -got,+want :%v ", diff)
			}
		})
	}
}
