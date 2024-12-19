package auth

import "context"

type LogoutUsecase struct {
	tokenAuthenticator           TokenAuthenticator
	tokenAuthenticatorRepository TokenAuthenticatorRepository
}

func NewLogoutUsecase(
	tokenAuthenticator TokenAuthenticator,
	tokenAuthenticatorRepository TokenAuthenticatorRepository,
) *LogoutUsecase {
	return &LogoutUsecase{
		tokenAuthenticator:           tokenAuthenticator,
		tokenAuthenticatorRepository: tokenAuthenticatorRepository,
	}
}

func (lu *LogoutUsecase) Run(ctx context.Context, input LogoutUsecaseInputDTO) error {
	if err := lu.tokenAuthenticatorRepository.Delete(ctx, input.ID); err != nil {
		return err
	}
	return nil
}
