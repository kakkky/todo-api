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
