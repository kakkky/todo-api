package auth

import (
	"context"

	"github.com/kakkky/app/domain/errors"
)

type AuthorizationUsecase struct {
	jwtAuthenticator           JwtAuthenticator
	jwtAuthenticatorRepository JwtAuthenticatorRepository
}

func NewAuthorizationUsecase(
	jwtAuthenticator JwtAuthenticator,
	jwtAuthenticatorRepository JwtAuthenticatorRepository,
) *AuthorizationUsecase {
	return &AuthorizationUsecase{
		jwtAuthenticator:           jwtAuthenticator,
		jwtAuthenticatorRepository: jwtAuthenticatorRepository,
	}
}

func (au *AuthorizationUsecase) Run(ctx context.Context, input AuthorizationInputDTO) (
	*AuthorizationOutputDTO,
	error,
) {
	// 公開鍵で署名済みトークンを検証する
	userID, jti, err := au.jwtAuthenticator.VerifyJwtToken(input.SignedToken)
	if err != nil {
		return nil, err
	}
	// KVS から保存された jti を取得
	// ログアウトしていた場合は、nilが返る
	jtiFromKVS, err := au.jwtAuthenticatorRepository.Load(ctx, userID)
	if err != nil {
		return nil, err
	}
	// jti が一致しない場合はエラー
	// ログアウトしている場合は、ここでエラーとなる
	if jti != jtiFromKVS {
		return nil, errors.New("invalid JWT ID")
	}
	return &AuthorizationOutputDTO{
		UserID: userID,
	}, nil
}
