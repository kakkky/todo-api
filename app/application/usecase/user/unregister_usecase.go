package user

import (
	"context"

	"github.com/kakkky/app/domain/user"
)

type unregisterUsecase struct {
	userRepository user.UserRepository
}

func NewUnregisterUsecase(
	userRepository user.UserRepository,
) *unregisterUsecase {
	return &unregisterUsecase{
		userRepository: userRepository,
	}
}

func (uu *unregisterUsecase) Run(ctx context.Context, input UnregisterUsecaseInputDTO) error {
	// 存在しているユーザーしか削除できない
	u, err := uu.userRepository.FindById(ctx, input.ID)
	// エラーかユーザーがnilの場合はエラー
	if err != nil || u == nil {
		return err
	}
	if err := uu.userRepository.Delete(ctx, u); err != nil {
		return err
	}
	return nil
}
