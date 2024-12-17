package user

import (
	"context"

	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/app/domain/user"
)

type RegisterUsecase struct {
	userRepository    user.UserRepository
	userDomainService user.UserDomainService
}

func NewRegisterUsecase(
	userRepository user.UserRepository,
	userDomainService user.UserDomainService,
) *RegisterUsecase {
	return &RegisterUsecase{
		userRepository:    userRepository,
		userDomainService: userDomainService,
	}
}

func (ru *RegisterUsecase) Run(ctx context.Context, input RegisterUsecaseInputDTO) (*RegisterUsecaseOutputDTO, error) {
	// userインスタンスを生成
	u, err := user.NewUser(
		input.Email,
		input.Name,
		input.Password,
	)
	if err != nil {
		return nil, err
	}
	//userがすでに登録済みでないか？
	ok, err := ru.userDomainService.IsExists(ctx, u.GetEmail())
	if err != nil {
		return nil, err
	}
	// emailがすでに存在していた場合
	if ok {
		return nil, errors.ErrAlreadyRegisterd
	}
	if err := ru.userRepository.Save(ctx, u); err != nil {
		return nil, err
	}
	// DTOに詰め替えて返す
	return &RegisterUsecaseOutputDTO{
		ID:    u.GetID(),
		Name:  u.GetName(),
		Email: u.GetEmail().Value(),
	}, nil
}
