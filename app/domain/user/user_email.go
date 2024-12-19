package user

import (
	"net/mail"

	"github.com/kakkky/app/domain/errors"
)

type Email struct {
	value string
}

func NewEmail(value string) (Email, error) {
	// バリデーション
	if _, err := mail.ParseAddress(value); err != nil {
		return Email{}, errors.ErrInvalidEmail
	}
	return Email{value: value}, nil
}

func reconstructEmail(value string) Email {
	return Email{value: value}
}

func (e Email) Value() string {
	return e.value
}
