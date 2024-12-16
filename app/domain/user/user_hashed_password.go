package user

import (
	"unicode/utf8"

	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/pkg/hash"
)

type HashedPassword struct {
	value string
}

const (
	minPasswordLength = 6
)

func newHashedPassword(value string) (HashedPassword, error) {
	// バリデーション
	if minPasswordLength >= utf8.RuneCountInString(value) {
		return HashedPassword{}, errors.ErrPasswordTooShort
	}
	// ハッシュ化する
	hashed, err := hash.Hash(value)
	if err != nil {
		return HashedPassword{}, err
	}

	return HashedPassword{value: hashed}, nil
}

func reconstructHashedPassword(value string) HashedPassword {
	return HashedPassword{value: value}
}

func (hp HashedPassword) Value() string {
	return hp.value
}

// ハッシュ化されたパスワードと比較
// 集約ルートUserから呼び出す
func (p HashedPassword) compare(target string) bool {
	return hash.Compare(p.value, target)
}
