package user

import (
	"unicode/utf8"

	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/pkg/hash"
)

type hashedPassword struct {
	value string
}

const (
	minPasswordLength = 6
)

func newHashedPassword(value string) (hashedPassword, error) {
	// バリデーション
	if minPasswordLength >= utf8.RuneCountInString(value) {
		return hashedPassword{}, errors.ErrPasswordTooShort
	}
	// ハッシュ化する
	hashed, err := hash.Hash(value)
	if err != nil {
		return hashedPassword{}, err
	}

	return hashedPassword{value: hashed}, nil
}

func reconstructHashedPassword(value string) hashedPassword {
	return hashedPassword{value: value}
}

func (hp hashedPassword) Value() string {
	return hp.value
}

// ハッシュ化されたパスワードと比較
func (p hashedPassword) Compare(target string) bool {
	return hash.Compare(p.value, target)
}
