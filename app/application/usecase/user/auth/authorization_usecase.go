package auth

import (
	"context"

	"github.com/kakkky/app/domain/errors"
)

type AuthorizationUsecase struct {
	tokenAuthenticator           TokenAuthenticator
	tokenAuthenticatorRepository TokenAuthenticatorRepository
}

func NewAuthorizationUsecase(
	tokenAuthenticator TokenAuthenticator,
	tokenAuthenticatorRepository TokenAuthenticatorRepository,
) *AuthorizationUsecase {
	return &AuthorizationUsecase{
		tokenAuthenticator:           tokenAuthenticator,
		tokenAuthenticatorRepository: tokenAuthenticatorRepository,
	}
}

func (au *AuthorizationUsecase) Run(ctx context.Context, input AuthorizationInputDTO) (
	*AuthorizationOutputDTO,
	error,
) {
	// 公開鍵で署名済みトークンを検証する
	// 解読されたトークンが返る
	token, err := au.tokenAuthenticator.VerifyToken(input.SignedToken)
	if err != nil {
		return nil, err
	}
	// トークンの有効期限を検証
	if err := au.tokenAuthenticator.VerifyExpiresAt(token); err != nil {
		return nil, err
	}
	// JWT クレームから情報を取得
	jti, err := au.tokenAuthenticator.GetJWTIDFromClaim(token)
	if err != nil {
		return nil, err
	}
	userID, err := au.tokenAuthenticator.GetSubFromClaim(token)
	if err != nil {
		return nil, err
	}
	// KVS から保存された jti を取得
	// ログアウトしていた場合は、nilが返る
	jtiFromKVS, err := au.tokenAuthenticatorRepository.Load(ctx, userID)
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
