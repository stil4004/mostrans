package main

import (
	"log"
	"net/http"
	"service/pkg/chat"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Response struct {
	Data string `json:"data"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/api/admin/login", func(c *fiber.Ctx) error {
		log.Println(string(c.Body()))
		var resp Response = Response{
			Data: "все неправильно, чмо",
		}
		return c.Status(200).JSON(resp)
	})
	app.Post("/chat", func(c *fiber.Ctx) error {
		log.Println(string(c.Body()))
		var resp Response = Response{
			Data: "хуйню сказала",
		}
		return c.Status(200).JSON(resp)
	})
	go func() {
		for {
			log.Println("running...")
			time.Sleep(5 * time.Second)
		}
	}()
	
	app.Listen("0.0.0.0:12060")
}

func main_old() {

	http.HandleFunc("/hell", chat.SocketHandler)
	http.HandleFunc("/", chat.Helloer)
	log.Fatal(http.ListenAndServe("localhost:12060", nil))
}
