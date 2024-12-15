package user

import "context"

//go:generate mockgen -package=user -source=./interface_user_repository.go -destination=./mock_user_repository.go
type UserRepository interface {
	Save(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email Email) (*User, error)
	FindById(ctx context.Context, id string) (*User, error)
	FetchAllUsers(ctx context.Context) (Users, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, user *User) error
}
