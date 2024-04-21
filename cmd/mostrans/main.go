package main

import (
	"log"
	"net/http"
	"service/config"
	"service/internal/server"
	"service/pkg/chat"
)

type Response struct {
	Data string `json:"data"`
}

func main() {
	viperInstance, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot load config. Error: {%s}", err.Error())
	}

	cfg, err := config.ParseConfig(viperInstance)
	if err != nil {
		log.Fatalf("Cannot parse config. Error: {%s}", err.Error())
	}

	s := server.NewServer(cfg)
	if err = s.Run(); err != nil {
		log.Println("err run:", err)
	}
}

