package auth

type AuthorizationInputDTO struct {
	JwtToken string
}

type AuthorizationOutputDTO struct {
	UserID string
}
