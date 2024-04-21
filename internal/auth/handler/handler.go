package handler

import (
	"service/config"
	"service/internal/auth"
)

type AuthHandler struct {
	cfg    *config.Config
	authUC auth.UseCase
}

func NewAuthHandler(cfg *config.Config, authUC auth.UseCase) *AuthHandler {
	return &AuthHandler{
		cfg:    cfg,
		authUC: authUC,
	}
}
