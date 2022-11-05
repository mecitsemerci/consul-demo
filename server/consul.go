package server

import (
	consul "github.com/hashicorp/consul/api"
	"log"
	"time"
)

func (s *Server) registerConsul() {
	c, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		log.Fatalln("Consul client error ", err.Error())
	}
	s.ConsulAgent = c.Agent()

	serviceDef := &consul.AgentServiceRegistration{
		Name: s.Name,
		Check: &consul.AgentServiceCheck{
			TTL: s.TTL.String(),
		},
	}

	if err := s.ConsulAgent.ServiceRegister(serviceDef); err != nil {
		log.Fatalln("Consul registration error ", err.Error())
	}

}

func (s *Server) UpdateTTL(check func() (bool, error)) {
	ticker := time.NewTicker(s.TTL / 2)
	for range ticker.C {
		s.update(check)
	}
}

func (s *Server) update(check func() (bool, error)) {
	if s.ConsulAgent == nil {
		log.Println("Agent is nil")
		return
	}
	ok, err := check()
	if !ok {
		log.Printf("err=\"Check failed\" msg=\"%s\"", err.Error())
		if agentErr := s.ConsulAgent.FailTTL("service:"+s.Name, err.Error()); agentErr != nil {
			log.Print(agentErr)
		}
	} else {
		if agentErr := s.ConsulAgent.PassTTL("service:"+s.Name, "Redis is connected successfully"); agentErr != nil {
			log.Print(agentErr)
		}
	}
}
