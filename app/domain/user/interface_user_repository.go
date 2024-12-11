package user

import "context"

//go:generate mockgen -package=user -source=./interface_user_repository.go -destination=./mock_user_repository.go
type UserRepository interface {
	FindByEmail(ctx context.Context, email email) (*User, error)
}
