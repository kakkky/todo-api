package errors

import (
	"errors"
)

// ユーザー関連
var (
	ErrInvalidEmail     = errors.New("invalid email address")
	ErrPasswordTooShort = errors.New("password is too short")
)

// errors.Isをラップ
// パッケージ名の衝突を考慮
func Is(err, target error) bool {
	return errors.Is(err, target)
}
