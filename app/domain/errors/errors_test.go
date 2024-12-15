package errors

import (
	"fmt"
	"testing"
)

func TestErrors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		fn   func() error
		err  *ErrDomain
	}{
		{
			name: "ErrDomain型としても、個別のエラー型としても判別できる",
			fn: func() error {
				return ErrNotFoundUser
			},
			err: ErrNotFoundUser,
		},
		{
			name: "エラー型がラップされていても判別できる",
			fn: func() error {
				wrappedErr := fmt.Errorf("error occured:%w", ErrNotFoundUser)
				return fmt.Errorf("error occured:%w", wrappedErr)
			},
			err: ErrNotFoundUser,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.fn()
			if err != nil && !Is(err, tt.err) {
				t.Errorf("failed to identificate error")
			}
			if !IsDomainErr(err) {
				t.Errorf("failed to identificate ErrDomain")
			}
		})
	}
}
