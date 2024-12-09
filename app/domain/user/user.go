package user

import (
	"github.com/kakkky/pkg/ulid"
)

type User struct {
	id             string
	email          email
	name           string
	hashedPassword hashedPassword
}

// 新たなユーザーを作成する
func NewUser(
	email string,
	name string,
	password string,
) (*User, error) {
	validatedEmail, err := newEmail(email)
	if err != nil {
		return nil, err
	}
	HashedPassword, err := newHashedPassword(password)
	if err != nil {
		return nil, err
	}
	return &User{
		id:             ulid.NewUlid(),
		email:          validatedEmail,
		name:           name,
		hashedPassword: HashedPassword,
	}, nil
}

// 既存のユーザーを返す
// リポジトリからのみ使用する
// インスタンスの再構成
func ReconstructUser(
	id string,
	email string,
	name string,
	hashedPassword string,
) *User {
	return &User{
		id:             id,
		email:          reconstructEmail(email),
		name:           name,
		hashedPassword: reconstructHashedPassword(hashedPassword),
	}
}
