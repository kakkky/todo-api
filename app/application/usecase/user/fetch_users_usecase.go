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

func (fuu *FetchUsersUsecase) Run(ctx context.Context) ([]*FetchUserUsecaseOutputDTO, error) {
	us, err := fuu.userRepository.FetchAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	outputs := make([]*FetchUserUsecaseOutputDTO, 0, len(us))
	for _, u := range us {
		outputs = append(outputs, &FetchUserUsecaseOutputDTO{ID: u.GetID(), Name: u.GetName()})
	}
	return outputs, nil
}
