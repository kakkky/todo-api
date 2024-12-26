package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kakkky/app/domain/errors"
)

func TestNewUser(t *testing.T) {
	t.Parallel()
	type args struct {
		name     string
		email    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		errType error
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				name:     "test",
				email:    "example@test.com",
				password: "password",
			},
			want: &User{
				email: Email{value: "example@test.com"},
				name:  "test",
			},
			wantErr: false,
		},
		{
			name: "準正常系：emailアドレスが不正",
			args: args{
				email: "test.com",
			},
			errType: errors.ErrInvalidEmail,
			wantErr: true,
		},
		{
			name: "準正常系：パスワードが短い(4文字)",
			args: args{
				password: "test",
			},
			errType: errors.ErrPasswordTooShort,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewUser(tt.args.email, tt.args.name, tt.args.password)
			if (err != nil) != tt.wantErr && errors.Is(err, tt.errType) {
				t.Fatalf("NewUser() error=%v,but wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(User{}, Email{}, HashedPassword{}), cmpopts.IgnoreFields(User{}, "id", "hashedPassword")); diff != "" {
				t.Errorf("NewProduct() -got,+want :%v ", diff)
			}
		})
	}
}
func TestUser_UpdateUser(t *testing.T) {
	user, _ := NewUser(
		"email@test.com",
		"testuser",
		"password",
	)
	t.Parallel()
	type args struct {
		email string
		name  string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		errType error
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				email: "example@test.com",
				name:  "test",
			},
			want: &User{
				id:             "1",
				email:          Email{value: "example@test.com"},
				name:           "test",
				hashedPassword: HashedPassword{value: "hashedPassword"},
			},
			wantErr: false,
		},
		{
			name: "準正常系：emailアドレスが不正",
			args: args{
				email: "invalid-email",
				name:  "test",
			},
			errType: errors.ErrInvalidEmail,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := user.UpdateUser(tt.args.email, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Fatalf("UpdateUser() error=%v, wantErr=%v", err, tt.wantErr)
			}
			if err != nil && !errors.Is(err, tt.errType) {
				t.Fatalf("UpdateUser() error type=%v, wantErr type=%v", err, tt.errType)
			}
			if diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(User{}, Email{}, HashedPassword{}), cmpopts.IgnoreFields(User{}, "id", "hashedPassword")); diff != "" {
				t.Errorf("UpdateUser() -got,+want :%v", diff)
			}
		})
	}
}

func TestCompairePassword(t *testing.T) {
	t.Parallel()
	// パスワードを持ったユーザーを用意
	user, _ := NewUser(
		"email@test.com",
		"testuser",
		"password",
	)
	tests := []struct {
		name    string
		plain   string
		wantErr bool
	}{
		{
			name:    "正常形：パスワードの比較で成功する",
			plain:   "password",
			wantErr: false,
		},
		{
			name:    "準正常形：パスワードの比較で成功する",
			plain:   "incorrect",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := user.ComparePassword(tt.plain)
			if (err != nil) != tt.wantErr && errors.Is(err, errors.ErrPasswordMismatch) {
				t.Errorf("ComparePassword() = %v,want %v", err, tt.wantErr)
			}

		})
	}
}
