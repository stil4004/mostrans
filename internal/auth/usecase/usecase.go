package usecase

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"service/internal/auth"
)

const salt = "siojeojioj@OiJ"

type AuthUseCase struct {
	repo auth.Repository
}

func NewAuthUseCase(repo auth.Repository) auth.UseCase {
	return &AuthUseCase{
		repo: repo,
	}
}

func (a *AuthUseCase) LogIn(ctx context.Context, req auth.LogInRequest) (auth.LogInResponse, error) {
	resp, err := a.repo.CheckLogIn(ctx, auth.CheckLogInRequest{
		NickName: req.NickName,
		Password: req.Password,
	})
	if err != nil {
		return auth.LogInResponse{
			Authenticated: false,
		}, err
	}

	if !resp.Authenticated {
		return auth.LogInResponse{
			Authenticated: false,
		}, errors.New("wrong password")
	}
	// TODO
	// token := jwt.NewWithClaims(
	// 	jwt.SigningMethodHS256,
	// 	&jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
	// 		IssuedAt:  time.Now().Unix(),

	// 	},
	// )

	return auth.LogInResponse{
		Authenticated: false,
	}, nil
}

func generatePasswordHashJWT(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
