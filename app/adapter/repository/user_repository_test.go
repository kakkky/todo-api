package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/app/domain/user"
	testhelper "github.com/kakkky/app/infrastructure/db/testhelper"
)

func TestUserRepository_Save(t *testing.T) {
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
	userRepository := NewUserRepository()
	user1, _ := user.NewUser(
		"user1@test.com",
		"user1",
		"password",
	)
	noExistentUser, _ := user.NewUser(
		"noexistent@test.com",
		"noexistent",
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
				email: noExistentUser.GetEmail(),
			},
			errType: errors.ErrNotFoundUser,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// user1のみがDBに保存されている
			testhelper.SetupFixtures(t, "testdata/fixtures/users.yml")
			ctx := context.Background()
			got, err := userRepository.FindByEmail(ctx, tt.args.email)
			// 期待されるエラータイプが設定されている場合はそれも検証
			if (err != nil) != tt.wantErr && tt.errType != nil {
				t.Errorf("userRepository.FindByEmail() =error:%v, want errType:%v", err, tt.errType)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FindByEmail() =error:%v, wantErr:%v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(user.User{}, user.Email{}, user.HashedPassword{}), cmpopts.IgnoreFields(user.User{}, "id", "hashedPassword")); diff != "" {
				t.Errorf("userRepository.FindByEmail() -got,+want :%v ", diff)
			}
		})
	}
}
func TestUserRepository_FindById(t *testing.T) {
	userRepository := NewUserRepository()
	user1 := user.ReconstructUser(
		"1",
		"user1@test.com",
		"user1",
		"password",
	)
	noExistentUser := user.ReconstructUser(
		"noExistent",
		"",
		"",
		"",
	)
	type args struct {
		id string
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
				id: user1.GetID(),
			},
			want:    user1,
			wantErr: false,
		},
		{
			name: "準正常系：ユーザーが見つからなければErrNotFoundUserが返ってくる",
			args: args{
				id: noExistentUser.GetID(),
			},
			errType: errors.ErrNotFoundUser,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// user1のみがDBに保存されている
			testhelper.SetupFixtures(t, "testdata/fixtures/users.yml")
			ctx := context.Background()
			got, err := userRepository.FindById(ctx, tt.args.id)
			// 期待されるエラータイプが設定されている場合はそれも検証
			if (err != nil) != tt.wantErr && tt.errType != nil {
				t.Errorf("userRepository.FindById() =error:%v, want errType:%v", err, tt.errType)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FindById() =error:%v, wantErr:%v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(user.User{}, user.Email{}, user.HashedPassword{}), cmpopts.IgnoreFields(user.User{}, "id", "hashedPassword")); diff != "" {
				t.Errorf("userRepository.FindById() -got,+want :%v ", diff)
			}
		})
	}
}

func TestUserRepository_FetchAllUsers(t *testing.T) {
	userRepository := NewUserRepository()
	var users user.Users
	for i := 0; i < 3; i++ {
		u, _ := user.NewUser(
			fmt.Sprintf("user%d@test.com", i+1),
			fmt.Sprintf("user%d", i+1),
			"password",
		)
		users = append(users, u)
	}
	tests := []struct {
		name    string
		want    user.Users
		wantErr bool
	}{
		{
			name:    "正常系：全てのユーザーを取得できる",
			want:    users,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			testhelper.SetupFixtures(t, "testdata/fixtures/users.yml")
			ctx := context.Background()
			got, err := userRepository.FetchAllUsers(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FetchAllUsers() =error:%v, wantErr:%v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(user.User{}, user.Email{}, user.HashedPassword{}), cmpopts.IgnoreFields(user.User{}, "id", "hashedPassword")); diff != "" {
				t.Errorf("userRepository.FetchAllUsers() -got,+want :%v ", diff)
			}
		})
	}
}

func TestUserRepository_Update(t *testing.T) {
	userRepository := NewUserRepository()
	// 更新情報を詰め替えて再構成したユーザー
	updatingUser := user.ReconstructUser(
		"1",
		"updated@test.com", // 更新
		"updatedUser",      // 更新
		"password",
	)
	updatingNameUser := user.ReconstructUser(
		"1",
		"user1@test.com",
		"nameUpdatedUser", //更新
		"password",
	)
	type args struct {
		user *user.User
	}
	tests := []struct {
		name    string
		args    args
		want    *user.User
		wantErr bool
	}{
		{
			name: "正常系: emailとnameの両方を更新する",
			args: args{
				user: updatingUser,
			},
			want:    updatingUser,
			wantErr: false,
		},
		{
			name: "正常系: nameのみを更新する",
			args: args{
				user: updatingNameUser,
			},
			want:    updatingNameUser,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			testhelper.SetupFixtures(t, "testdata/fixtures/users.yml")
			ctx := context.Background()
			if err := userRepository.Update(ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userRepository.Update()=error:%v, wantErr:%v", err, tt.wantErr)
			}
			got, err := userRepository.FindByEmail(ctx, tt.args.user.GetEmail())
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FindByEmail()=error:%v, wantErr:%v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(user.User{}, user.Email{}, user.HashedPassword{}), cmpopts.IgnoreFields(user.User{}, "id", "hashedPassword")); diff != "" {
				t.Errorf("userRepository.FindByEmail() -got,+want :%v ", diff)
			}
		})
	}
}
func TestUserRepository_Delete(t *testing.T) {
	userRepository := NewUserRepository()
	deletingUser := user.ReconstructUser(
		"1",
		"user1@test.com",
		"user1",
		"password",
	)
	type args struct {
		user *user.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系: 指定のユーザーを削除できる",
			args: args{
				user: deletingUser,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			testhelper.SetupFixtures(t, "testdata/fixtures/users.yml")
			ctx := context.Background()
			if err := userRepository.Delete(ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userRepository.Delete()=error:%v, wantErr:%v", err, tt.wantErr)
			}
			_, err := userRepository.FindByEmail(ctx, deletingUser.GetEmail())
			// エラーがあればそれはErrNotFoundUserであるべき
			if err != nil && !errors.Is(err, errors.ErrNotFoundUser) {
				t.Errorf("userRepository.FindByEmail()=error:%v, wantErr:%v", err, tt.wantErr)
			}
		})
	}
}
