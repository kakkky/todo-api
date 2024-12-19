package repository

import (
	"context"
	"testing"
	"time"
)

func TestTokenAuthenticatorRepository_Save_And_Load(t *testing.T) {
	tokenAuthenticatorRepository := NewTokenAuthenticatorRepository()
	type args struct {
		user_id  string
		jwt_id   string
		duration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "正常系: user_id:jwt_idのペアでredisに保存できる",
			args: args{
				user_id:  "1",
				jwt_id:   "jwt",
				duration: time.Duration(1 * time.Minute),
			},
			want:    "jwt",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			if err := tokenAuthenticatorRepository.Save(ctx, tt.args.duration, tt.args.user_id, tt.args.jwt_id); (err != nil) != tt.wantErr {
				t.Errorf("tokenAuthenticatorRepository.Save() error=%v,wantErr %v", err, tt.wantErr)
			}
			got, err := tokenAuthenticatorRepository.Load(ctx, tt.args.user_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("tokenAuthenticatorRepository.Load() error=%v,wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("tokenAuthenticatorRepository.Load() got %v, want %v", got, tt.want)
			}
		})
	}
}
