package user

import (
	"github.com/kakkky/pkg/ulid"
)

type User struct {
	id       string
	email    email
	name     string
	password password
}

// 新たなユーザーを作成する
func NewUser(
	id string,
	email string,
	name string,
	password string,
) (*User, error) {
	return newUser(
		ulid.NewUlid(),
		email,
		name,
		password,
	)
}

// 既存のユーザーを返す
func ReconstructUser(
	id string,
	email string,
	name string,
	password string,
) (*User, error) {
	return newUser(
		id,
		email,
		name,
		password,
	)
}

// コンストラクタ
func newUser(
	id string,
	email string,
	name string,
	password string,
) (*User, error) {
	validatedEmail, err := newEmail(email)
	if err != nil {
		return nil, err
	}
	validatedPassword, err := newPassword(password)
	if err != nil {
		return nil, err
	}
	return &User{
		id:       id,
		email:    validatedEmail,
		name:     name,
		password: validatedPassword,
	}, nil
}

// ゲッターメソッド
func (u *User) ID() string {
	return u.id
}
func (u *User) Email() string {
	return u.email.value
}
func (u *User) Name() string {
	return u.name
}
func (u *User) Password() string {
	return u.password.value
}
