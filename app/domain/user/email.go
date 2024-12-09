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
	if err, _ := mail.ParseAddress(value); err != nil {
		return email{}, errors.ErrInvalidEmail
	}
	return email{value: value}, nil
}
