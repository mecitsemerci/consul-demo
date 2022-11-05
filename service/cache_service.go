package service

import (
	"github.com/go-redis/redis"
)

type Service struct {
	RedisClient redis.UniversalClient
}

func New(addrs []string) (*Service, error) {
	s := new(Service)

	s.RedisClient = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: addrs,
	})

	if ok, err := s.Check(); !ok {
		return nil, err
	}

	return s, nil
}

func (s *Service) Check() (bool, error) {
	if _, err := s.RedisClient.Ping().Result(); err != nil {
		return false, err
	}
	return true, nil
}
