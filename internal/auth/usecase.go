package auth

import "context"

type UseCase interface {
	LogIn(ctx context.Context, req LogInRequest) (LogInResponse, error)
}
