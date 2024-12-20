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
			sut := NewJWTAuthenticator()

			// トークンを生成する
			token := sut.GenerateToken(tt.args.sub, tt.args.jti)
			// トークンに署名し、署名済みトークンを生成
			signedToken, err := sut.SignToken(token)
			if err != nil {
				t.Errorf("SignToken error = %v", err)
			}

			// 署名済みトークンを解析する
			verifiedToken, err := sut.VerifyToken(signedToken)
			if err != nil {
				t.Errorf("VerifyToken error = %v", err)
			}

			if err := sut.VerifyExpiresAt(verifiedToken); err != nil {
				t.Errorf("VerifyExpiresAt error = %v", err)
			}

			gotJti, err := sut.GetJWTIDFromClaim(verifiedToken)
			if err != nil {
				t.Errorf("GetJWTIDFromClaim error = %v", err)
			}
			if gotJti != tt.args.jti {
				t.Errorf("mismatch jti: expected %v, got %v", tt.args.jti, gotJti)
			}

			gotSub, err := sut.GetSubFromClaim(verifiedToken)
			if err != nil {
				t.Errorf("GetSubFromClaim error = %v", err)
			}
			if gotSub != tt.args.sub {
				t.Errorf("mismatch sub: expected %v, got %v", tt.args.sub, gotSub)
			}
		})
	}
}
