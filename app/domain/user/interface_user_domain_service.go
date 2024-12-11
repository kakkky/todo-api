package user

import "context"

type UserDomainService interface {
	IsExists(ctx context.Context, email email) (bool, error)
}
