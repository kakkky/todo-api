package hash

import (
	"testing"
)

func TestHash(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "valid password",
			password: "testPassword",
			wantErr:  false,
		},
		{
			name:     "empty password",
			password: "",
			wantErr:  false,
		},
		{
			name:     "long password",
			password: "verylongpassword123456",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ハッシュ化
			hashedPassword, err := Hash(tt.password)

			// エラーの検証
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error: %v, got: %v", tt.wantErr, err)
			}

			// エラーがない場合、ハッシュ化されたパスワードが空でないことを確認
			if !tt.wantErr && hashedPassword == "" {
				t.Fatal("hashed password should not be empty")
			}
		})
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		name           string
		password       string
		targetPassword string
		want           bool
	}{
		{
			name:           "correct password match",
			password:       "testPassword",
			targetPassword: "testPassword",
			want:           true,
		},
		{
			name:           "incorrect password match",
			password:       "testPassword",
			targetPassword: "wrongPassword",
			want:           false,
		},
		{
			name:           "empty password match",
			password:       "",
			targetPassword: "",
			want:           true,
		},
		{
			name:           "empty password mismatch",
			password:       "testPassword",
			targetPassword: "",
			want:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ハッシュ化
			hashedPassword, err := Hash(tt.password)
			if err != nil {
				t.Fatalf("expected no error during hashing, got %v", err)
			}

			// 比較結果の検証
			result := Compare(hashedPassword, tt.targetPassword)
			if result != tt.want {
				t.Errorf("expected result: %v, got: %v", tt.want, result)
			}
		})
	}
}
