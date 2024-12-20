package auth

type AuthorizationInputDTO struct {
	SignedToken string
}

type AuthorizationOutputDTO struct {
	UserID string
}
