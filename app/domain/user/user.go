package user

import (
	"github.com/kakkky/pkg/ulid"
)

type User struct {
	id             string
	email          Email
	name           string
	hashedPassword HashedPassword
}

// ファクトリー関数
func NewUser(
	email string,
	name string,
	password string,
) (*User, error) {
	validatedEmail, err := NewEmail(email)
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

// ユーザーオブジェクト更新
func (u *User) UpdateUser(
	email string,
	name string,
) (*User, error) {
	validatedEmail, err := NewEmail(email)
	if err != nil {
		return nil, err
	}
	return &User{
		id:             u.id,
		email:          validatedEmail,
		name:           name,
		hashedPassword: u.hashedPassword,
	}, nil
}

// 値のゲッターメソッド
func (u *User) GetID() string {
	return u.id
}
func (u *User) GetName() string {
	return u.name
}
func (u *User) GetEmail() Email {
	return u.email
}
func (u *User) GetHashedPassword() HashedPassword {
	return u.hashedPassword
}

// パスワードを比較する
func (u *User) ComparePassword(plainPassword string) error {
	return u.hashedPassword.compare(plainPassword)
}
