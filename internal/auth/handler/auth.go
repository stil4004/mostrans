package handler

import (
	"net/http"
	"service/internal/auth"

	"github.com/gofiber/fiber/v2"
)

func (h *AuthHandler) LogIn() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params auth.CheckLogInRequest

		if err := ctx.BodyParser(&params); err != nil {
			return err
		}

		resp, err := h.authUC.LogIn(ctx.Context())
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(map[string]interface{}{
				"error": err.Error(),
				"auth":  resp.Authenticated,
			})
		}

		return ctx.Status(http.StatusOK).JSON(resp)
	}
}
