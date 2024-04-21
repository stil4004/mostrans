package handler

import (
	"service/config"
	"service/internal/chat"
)

type ChatHandler struct {
	cfg    *config.Config
	chatUC chat.UseCase
}

func NewAuthHandler(cfg *config.Config, chatUC chat.UseCase) *ChatHandler {
	return &ChatHandler{
		cfg:    cfg,
		chatUC: chatUC,
	}
}
