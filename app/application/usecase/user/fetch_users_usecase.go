package user

import (
	"context"

	"github.com/kakkky/app/domain/user"
)

type FetchUsersUsecase struct {
	userRepository user.UserRepository
}

func NewFetchUsersUsecase(
	userRepository user.UserRepository,
) *FetchUsersUsecase {
	return &FetchUsersUsecase{
		userRepository: userRepository,
	}
}

func (luu *FetchUsersUsecase) Run(ctx context.Context) ([]*FetchUsersUsecaseOutputDTO, error) {
	us, err := luu.userRepository.FetchAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	outputs := make([]*FetchUsersUsecaseOutputDTO, 0, len(us))
	for _, u := range us {
		outputs = append(outputs, &FetchUsersUsecaseOutputDTO{ID: u.GetID(), Name: u.GetName()})
	}
	return outputs, nil
}
