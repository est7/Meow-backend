package interfaces

import (
	"github.com/redis/go-redis/v9"
)

// Service 定义了 Service 接口
// redis缓存在 service 层处理,db 在 repository 层处理
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
