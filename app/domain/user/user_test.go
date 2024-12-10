package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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
				email: email{value: "example@test.com"},
				name:  "test",
			},
			wantErr: false,
		},
		{
			name: "異常系：emailアドレスが不正",
			args: args{
				email: "test.com",
			},
			wantErr: true,
		},
		{
			name: "異常系：パスワードが短い(4文字)",
			args: args{
				password: "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewUser(tt.args.email, tt.args.name, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewUser() error=%v,but wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(User{}, email{}, hashedPassword{}), cmpopts.IgnoreFields(User{}, "id", "hashedPassword")); diff != "" {
				t.Errorf("NewProduct() -got,+want :%v ", diff)
			}
		})
	}
}
