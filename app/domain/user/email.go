package user

import (
	"net/mail"

	"github.com/kakkky/app/domain/errors"
)

type email struct {
	value string
}

func newEmail(value string) (email, error) {
	// バリデーション
	if _, err := mail.ParseAddress(value); err != nil {
		return email{}, errors.ErrInvalidEmail
	}
	return email{value: value}, nil
}

func reconstructEmail(value string) email {
	return email{value: value}
}
