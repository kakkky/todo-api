package errors

import (
	"errors"
)

// ユーザー関連
var (
	ErrAlreadyRegisterd = errors.New("you have been already registerd")
	ErrInvalidEmail     = errors.New("invalid email address")
	ErrPasswordTooShort = errors.New("password is too short")
	ErrNotFoundUser     = errors.New("user not found")
)

// errors.Isをラップ
// パッケージ名の衝突を考慮
func Is(err, target error) bool {
	return errors.Is(err, target)
}

func New(message string) error {
	return errors.New(message)
}
