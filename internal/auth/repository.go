package auth

import "context"

type Repository interface {
	CheckLogIn(ctx context.Context, req CheckLogInRequest) (CheckLogInResponse, error)
	GetUserByLogin(ctx context.Context, req GetUserByLoginRequest) (GetUserByLoginResponse, error)
}
