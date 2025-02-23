package user

import (
	"context"

	"github.com/kakkky/app/domain/user"
)

type FetchUserUsecase struct {
	userRepository user.UserRepository
}

func NewFetchUserUsecase(
	userRepository user.UserRepository,
) *FetchUserUsecase {
	return &FetchUserUsecase{
		userRepository: userRepository,
	}
}

func (fuu *FetchUserUsecase) Run(ctx context.Context, input FetchUserUsecaseInputDTO) (*FetchUserUsecaseOutputDTO, error) {
	u, err := fuu.userRepository.FindById(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	output := &FetchUserUsecaseOutputDTO{
		ID:   u.GetID(),
		Name: u.GetName(),
	}
	return output, nil
}
