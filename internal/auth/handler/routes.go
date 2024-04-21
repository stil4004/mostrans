package handler

import (
	"github.com/gofiber/fiber/v2"
)

func MapAdminsRoutes(router fiber.Router, a *AuthHandler) {
	rg := router.Group("/admin")
	rg.Post("/sign_in", a.LogIn())

}
