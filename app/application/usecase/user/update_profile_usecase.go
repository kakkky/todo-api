package user

import (
	"context"

	"github.com/kakkky/app/domain/user"
)

type UpdateProfileUsecase struct {
	userRepository user.UserRepository
}

func NewUpdateProfileUsecase(
	userRepository user.UserRepository,
) *UpdateProfileUsecase {
	return &UpdateProfileUsecase{
		userRepository: userRepository,
	}
}

func (epu *UpdateProfileUsecase) Run(ctx context.Context, input UpdateProfileUsecaseInputDTO) (
	*UpdateProfileUsecaseOutputDTO, error,
) {
	// 存在しているユーザーしか編集できない
	u, err := epu.userRepository.FindById(ctx, input.ID)
	// エラーかユーザーがnilの場合はエラー
	if err != nil || u == nil {
		return nil, err
	}
	// 値が空の場合は既存情報を入れる
	if input.Name == "" {
		input.Name = u.GetName()
	}
	if input.Email == "" {
		input.Email = u.GetEmail().Value()
	}
	// input情報をもとに、更新情報を反映したインスタンスを作成
	updatedUser, err := user.UpdateUser(
		input.ID,
		input.Email,
		input.Name,
		u.GetHashedPassword().Value(), //パスワードはそのまま
	)
	if err != nil {
		return nil, err
	}
	if err := epu.userRepository.Update(ctx, u); err != nil {
		return nil, err
	}
	// 更新したオブジェクトをDTOに詰め替える
	return &UpdateProfileUsecaseOutputDTO{
		ID:    updatedUser.GetID(),
		Email: updatedUser.GetEmail().Value(),
		Name:  updatedUser.GetName(),
	}, nil
}
