package user

import "context"

//go:generate mockgen -package=user -source=./interface_user_domain_service.go -destination=./mock_user_domain_service.go
type UserDomainService interface {
	IsExists(ctx context.Context, email Email) (bool, error)
}
