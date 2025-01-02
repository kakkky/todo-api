package errors

import (
	"testing"
)

func TestErrors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		fn   func() error
		err  *ErrDomain
		want bool
	}{
		{
			name: "ErrDomain型としても、個別のエラー型としても判別できる",
			fn: func() error {
				return ErrNotFoundUser
			},
			err:  ErrNotFoundUser,
			want: true,
		},
		{
			name: "異なるエラー型オブジェクトは判別しない",
			fn: func() error {
				return New("other err obj")
			},
			err:  ErrNotFoundUser,
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.fn()
			if err != nil && Is(err, tt.err) != tt.want {
				t.Errorf("Is() want %v,but got%v", tt.want, Is(err, tt.err))
			}
			if IsDomainErr(err) != tt.want {
				t.Errorf("IsDomainErr() want %v,but got%v", tt.want, IsDomainErr(err))
			}
		})
	}
}
