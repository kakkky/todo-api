package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/kakkky/app/adapter/repository"
	"github.com/kakkky/app/infrastructure/kvs"
)

func TestTokenAuthenticatorRepository_Save_And_Load_And_Delete(t *testing.T) {
	redisCom, err := kvs.NewRedisCommander()
	if err != nil {
		t.Fatal(err)
	}
	tokenAuthenticatorRepository := repository.NewTokenAuthenticatorRepository(redisCom)
	type args struct {
		userID   string
		jwtID    string
		duration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "正常系: userID:jwtIDのペアでredisに保存できる",
			args: args{
				userID:   "1",
				jwtID:    "jwt",
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
			if err := tokenAuthenticatorRepository.Save(ctx, tt.args.duration, tt.args.userID, tt.args.jwtID); (err != nil) != tt.wantErr {
				t.Errorf("tokenAuthenticatorRepository.Save() error=%v,wantErr %v", err, tt.wantErr)
			}
			got, err := tokenAuthenticatorRepository.Load(ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("tokenAuthenticatorRepository.Load() error=%v,wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("tokenAuthenticatorRepository.Load() got %v, want %v", got, tt.want)
			}
			if err := tokenAuthenticatorRepository.Delete(ctx, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("tokenAuthenticatorRepository.Delete() error=%v,wantErr %v", err, tt.wantErr)
			}
		})
	}
}
