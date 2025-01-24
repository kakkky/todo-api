package auth

import "testing"

func TestJWTAuthenticator(t *testing.T) {
	type args struct {
		sub string
		jti string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "subとjtiを元に署名済みjwtトークンを生成し、解析してsubとjtiを取り出せる",
			args: args{
				sub: "userID",
				jti: "jti",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut := NewJwtAuthenticator()
			signedToken, err := sut.GenerateJwtToken(tt.args.sub, tt.args.jti)
			if err != nil {
				t.Errorf("GenerateJwtToken() error=%v", err)
			}
			gotSub, gotJti, err := sut.VerifyJwtToken(signedToken)
			if err != nil {
				t.Errorf("VerifyToken() error = %v", err)
				return
			}
			// 同値性の確認
			if gotSub != tt.args.sub {
				t.Errorf("VerifyToken() got = %v, want = %v", gotSub, tt.args.sub)
			}
			if gotJti != tt.args.jti {
				t.Errorf("VerifyToken() got = %v, want = %v", gotJti, tt.args.jti)
			}
		})
	}
}
