package user

import (
	"context"

	"github.com/kakkky/app/domain/user"
)

type EditProfileUsecase struct {
	userRepository user.UserRepository
}

func NewEditProfileUsecase(
	userRepository user.UserRepository,
) *EditProfileUsecase {
	return &EditProfileUsecase{
		userRepository: userRepository,
	}
}

func (epu *EditProfileUsecase) Run(ctx context.Context, input EditProfileUsecaseInputDTO) (
	*EditProfileUsecaseOutputDTO, error,
) {
	// 存在しているユーザーしか編集できない
	u, err := epu.userRepository.FindById(ctx, input.ID)
	// エラーかユーザーがnilの場合はエラー
	if err != nil || u == nil {
		return nil, err
	}
	// input情報をもとに、更新情報を反映したインスタンスを作成
	u = user.ReconstructUser(
		input.ID,
		input.Email,
		input.Name,
		"",
	)
	if err := epu.userRepository.Update(ctx, u); err != nil {
		return nil, err
	}
	return &EditProfileUsecaseOutputDTO{
		ID:    u.GetID(),
		Email: u.GetEmail().Value(),
		Name:  u.GetName(),
	}, nil
}
