package auth

import (
	"context"
	"time"

	"github.com/kakkky/app/domain/user"
	"github.com/kakkky/pkg/ulid"
)

type LoginUsecase struct {
	userRepository             user.UserRepository
	jwtAuthenticatorRepository JwtAuthenticatorRepository
	jwtAuthenticator           JwtAuthenticator
}

func NewLoginUsecase(
	userRepository user.UserRepository,
	jwtAuthenticatorRepository JwtAuthenticatorRepository,
	jwtAuthenticator JwtAuthenticator,
) *LoginUsecase {
	return &LoginUsecase{
		userRepository:             userRepository,
		jwtAuthenticatorRepository: jwtAuthenticatorRepository,
		jwtAuthenticator:           jwtAuthenticator,
	}
}

func (lu *LoginUsecase) Run(ctx context.Context, input LoginUsecaseInputDTO) (
	*LoginUsecaseOutputDTO, error,
) {
	// emailの検証
	email, err := user.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}
	// ユーザー検索
	u, err := lu.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	// パスワードは正しいか
	if err := u.ComparePassword(input.Password); err != nil {
		return nil, err
	}
	// トークンを生成
	jti := ulid.NewUlid() //JWTトークンを識別するID
	jwtToken, err := lu.jwtAuthenticator.GenerateJwtToken(u.GetID(), jti)
	if err != nil {
		return nil, err
	}
	// トークンをkvsに保存
	if err := lu.jwtAuthenticatorRepository.Save(ctx, time.Duration(24*time.Hour), u.GetID(), jti); err != nil {
		return nil, err
	}
	return &LoginUsecaseOutputDTO{
		JwtToken: jwtToken,
	}, nil
}
