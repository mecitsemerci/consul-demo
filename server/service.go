package server

import (
	"github.com/mecitsemerci/consul-demo/service"
	"log"
	"os"
)

func (s *Server) registerService() {
	rURL := os.Getenv("REDIS_URL")
	log.Print("Redis URL:" + rURL)
	if rURL == "" {
		log.Fatalln("Redis URL is empty")
	}
	addresses := []string{rURL}
	srv, err := service.New(addresses)
	if err != nil {
		log.Fatalln("Application startup error " + err.Error())
	}
	s.Service = srv
}
