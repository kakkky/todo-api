package auth

import (
	"context"
	"time"

	"github.com/kakkky/app/domain/user"
	"github.com/kakkky/pkg/ulid"
)

type LoginUsecase struct {
	userRepository               user.UserRepository
	tokenAuthenticatorRepository TokenAuthenticatorRepository
	tokenAuthenticator           TokenAuthenticator
}

func NewLoginUsecase(
	userRepository user.UserRepository,
	tokenAuthenticatorRepository TokenAuthenticatorRepository,
	tokenAuthenticator TokenAuthenticator,
) *LoginUsecase {
	return &LoginUsecase{
		userRepository:               userRepository,
		tokenAuthenticatorRepository: tokenAuthenticatorRepository,
		tokenAuthenticator:           tokenAuthenticator,
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
	jwtId := ulid.NewUlid() //JWTトークンを識別するID
	token := lu.tokenAuthenticator.GenerateToken(u.GetID(), jwtId)
	// トークンをkvsに保存
	if err := lu.tokenAuthenticatorRepository.Save(ctx, time.Duration(2*time.Hour), u.GetID(), jwtId); err != nil {
		return nil, err
	}
	// 署名済みトークンを発行
	signedToken, err := lu.tokenAuthenticator.SignToken(token)
	if err != nil {
		return nil, err
	}
	return &LoginUsecaseOutputDTO{
		SignedToken: signedToken,
	}, nil
}
