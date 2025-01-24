package auth

import "context"

type LogoutUsecase struct {
	jwtAuthenticator           JwtAuthenticator
	jwtAuthenticatorRepository JwtAuthenticatorRepository
}

func NewLogoutUsecase(
	jwtAuthenticator JwtAuthenticator,
	jwtAuthenticatorRepository JwtAuthenticatorRepository,
) *LogoutUsecase {
	return &LogoutUsecase{
		jwtAuthenticator:           jwtAuthenticator,
		jwtAuthenticatorRepository: jwtAuthenticatorRepository,
	}
}

func (lu *LogoutUsecase) Run(ctx context.Context, input LogoutUsecaseInputDTO) error {
	if err := lu.jwtAuthenticatorRepository.Delete(ctx, input.UserID); err != nil {
		return err
	}
	return nil
}
