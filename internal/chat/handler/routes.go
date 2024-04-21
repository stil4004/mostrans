package handler

import (
	"github.com/gofiber/fiber/v2"
)

func MapAuthRoutes(router fiber.Router, a *ChatHandler) {
	rg := router.Group("/chat")
	rg.Post("/send_message", a.ProcessMessage())

}
