package handler

import (
	"service/internal/auth"

	"honnef.co/go/tools/config"
)

type AuthHandler struct {
	// logger   *logger.ApiLogger
	cfg    *config.Config
	authUC auth.UseCase
	// loki     loki.UseCase
	// adminUC  admins.AdminCase
	// cache    cacheAdmins.Usecase
	// debateUC debate.DebCase
	// comUC    communication.ComCase
	// reqUtil  *utils.Request
	// chatUC   dnd_chat.DndChatCase
}

func NewAuthHandler(cfg *config.Config, authUC auth.UseCase) *AuthHandler {
	return &AuthHandler{
		cfg:    cfg,
		authUC: authUC,
	}
}
