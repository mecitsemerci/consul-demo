package server

import (
	"github.com/gin-gonic/gin"
	consul "github.com/hashicorp/consul/api"
	"github.com/mecitsemerci/consul-demo/goserver/service"
	"time"
)

type Server struct {
	Name        string
	TTL         time.Duration
	ConsulAgent *consul.Agent
	Engine      *gin.Engine
	Service     *service.Service
}

func New() Server {
	s := Server{
		Name: "cache-api",
		TTL:  10 * time.Second,
	}

	s.RegisterCacheService()
	s.RegisterRoutes()
	s.ConsulServiceRegister()

	go s.UpdateTTL(s.Service.Check)

	return s
}

func (s *Server) Run() error {
	return s.Engine.Run()
}
