package handler

import (
	"log"
	"net/http"
	"service/internal/auth"

	"github.com/gofiber/fiber/v2"
)

func (h *AuthHandler) LogIn() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params auth.CheckLogInRequest

		log.Println(string(ctx.Body()))
		
		if err := ctx.BodyParser(&params); err != nil {
			return err
		}

		resp, err := h.authUC.LogIn(ctx.Context(), auth.LogInRequest{
			NickName: params.NickName,
			Password: params.Password,
		})
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(map[string]interface{}{
				"error": err.Error(),
				"auth":  resp.Authenticated,
			})
		}

		return ctx.Status(http.StatusOK).JSON(resp)
	}
}
