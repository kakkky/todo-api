package user

import (
	"unicode/utf8"

	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/pkg/hash"
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
	// ハッシュ化する
	hashed, err := hash.Hash(value)
	if err != nil {
		return password{}, err
	}

	return password{value: string(hashed)}, nil
}

func reconstructPassword(hashedValue string) password {
	return password{value: hashedValue}
}

// ハッシュ化されたパスワードと比較
func (p password) Compare(target string) bool {
	return hash.Compare(p.value, target)
}
