package config

import "testing"

func TestInitConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		setEnv  string
		wantErr bool
	}{
		{
			name:    "正常系ーADDRESS環境変数を読み込める",
			want:    "0.0.0.0",
			setEnv:  "0.0.0.0",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("ADDRESS", tt.setEnv)
			err := InitConfig()
			got := config.Server.Address

			if (err != nil) != tt.wantErr {
				t.Errorf("want error : %v ,but : %v", tt.wantErr, err)
			}

			if got != tt.want {
				t.Errorf("want %v , but got : %v", tt.want, got)
			}
		})
	}

}
