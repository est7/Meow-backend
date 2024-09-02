package interfaces

import (
	"github.com/redis/go-redis/v9"
)

type Service interface {
	GetRepo() Repository
	GetRedis() *redis.Client
}

type BaseService struct {
	Repo  Repository
	Redis *redis.Client
}

func (s *BaseService) GetRepo() Repository {
	return s.Repo
}

func (s *BaseService) GetRedis() *redis.Client {
	return s.Redis
}

func NewService(repo Repository, redis *redis.Client) Service {
	return &BaseService{Repo: repo, Redis: redis}
}
