package server

import (
	"fmt"
	"log"
	"service/config"
	authHandler "service/internal/auth/handler"
	authRepo "service/internal/auth/repo"
	authUseCase "service/internal/auth/usecase"
	chatHandler "service/internal/chat/handler"
	chatRepo "service/internal/chat/repo"
	chatUsecase "service/internal/chat/usecase"

	"service/pkg/ai_client"
	"service/pkg/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	cfg   *config.Config
	fiber *fiber.App
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		fiber: fiber.New(fiber.Config{}),
		cfg:   cfg,
	}
}

func (s *Server) Run() error {
	if err := s.MapHandlers(s.fiber); err != nil {
		log.Println("can't handle ", err)
	}

	if err := s.fiber.Listen(fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port)); err != nil {
		log.Println("can't listen ", err)
	}

	return nil
}

func (s *Server) MapHandlers(app *fiber.App) error {
	db, err := storage.InitPsqlDB(s.cfg)
	if err != nil {
		return err
	}

	authRep := authRepo.NewPostgresRepository(db)
	authUC := authUseCase.NewAuthUseCase(authRep)
	authHandl := authHandler.NewAuthHandler(s.cfg, authUC)

	apiClientUC, err := ai_client.New(s.cfg)
	if err != nil {
		return err
	}

	chatRep := chatRepo.NewChatRepository(db)
	chatUC := chatUsecase.NewChatUseCase(chatRep, apiClientUC)
	chatHndl := chatHandler.NewAuthHandler(s.cfg, chatUC)
	app.Use(cors.New(cors.Config{}))

	chatUC.Cache()

	authHandler.MapAuthRoutes(app, authHandl)
	chatHandler.MapAuthRoutes(app, chatHndl)

	return nil
}
