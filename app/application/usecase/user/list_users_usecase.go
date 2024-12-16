package user

import (
	"context"

	"github.com/kakkky/app/domain/user"
)

type listUsersUsecase struct {
	userRepository user.UserRepository
}

func NewListUsersUsecase(
	userRepository user.UserRepository,
) *listUsersUsecase {
	return &listUsersUsecase{
		userRepository: userRepository,
	}
}

func (luu *listUsersUsecase) Run(ctx context.Context) ([]*ListUsersUsecaseOutputDTO, error) {
	us, err := luu.userRepository.FetchAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	outputs := make([]*ListUsersUsecaseOutputDTO, 0, len(us))
	for _, u := range us {
		outputs = append(outputs, &ListUsersUsecaseOutputDTO{ID: u.GetID(), Name: u.GetName()})
	}
	return outputs, nil
}
