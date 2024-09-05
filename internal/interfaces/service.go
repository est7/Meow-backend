package interfaces

import (
	"github.com/redis/go-redis/v9"
)

// BaseService 定义了 BaseService 接口
// redis缓存在 service 层处理,db 在 repository 层处理
type BaseService interface {
	GetRepo() Repository
	GetRedis() *redis.Client
}

type baseServiceImpl struct {
	Repo  Repository
	Redis *redis.Client
}

func (s *baseServiceImpl) GetRepo() Repository {
	return s.Repo
}

func (s *baseServiceImpl) GetRedis() *redis.Client {
	return s.Redis
}

func NewBaseService(repo Repository, redis *redis.Client) BaseService {
	return &baseServiceImpl{
		Repo:  repo,
		Redis: redis,
	}
}
