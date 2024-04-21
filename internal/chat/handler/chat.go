package handler

import (
	"log"
	"net/http"
	"service/internal/chat"

	"github.com/gofiber/fiber/v2"
)

func (h *ChatHandler) ProcessMessage() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params chat.ChatRequest

		// Token validation

		if err := ctx.BodyParser(&params); err != nil {
			return err
		}
		log.Println(params)

		// TODO handler
		// resp, err := h.authUC.LogIn(ctx.Context(), auth.LogInRequest{
		// 	NickName: params.NickName,
		// 	Password: params.Password,
		// })
		// if err != nil {
		// 	return ctx.Status(http.StatusUnauthorized).JSON(map[string]interface{}{
		// 		"error": err.Error(),
		// 		"auth":  resp.Authenticated,
		// 	})
		// }
		resp, err := h.chatUC.ProcessMessage(ctx.Context(), chat.ProcessMessageRequest{
			AIType:      params.AIType,
			MessageText: params.MessageText,
		})
		if err != nil {
			log.Println("Ошибка в чате ", params.MessageText, err)
			return ctx.Status(http.StatusOK).JSON(resp)
		}

		return ctx.Status(http.StatusOK).JSON(resp)
	}
}
