package user

import (
	"unicode/utf8"

	"github.com/kakkky/app/domain/errors"
)

type password struct {
	value string
}

const (
	minPasswordLength = 6
)

func newPassword(value string) (password, error) {
	// バリデーション
	if minPasswordLength >= utf8.RuneCountInString(value) {
		return password{}, errors.ErrPasswordTooShort
	}
	return password{value: value}, nil
}
