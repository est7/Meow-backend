package interfaces

import (
	"github.com/redis/go-redis/v9"
)

type ServiceFactory interface {
	CreateService(repo Repository, redis *redis.Client) Service
}

type BaseServiceFactory struct{}

func (f *BaseServiceFactory) CreateService(repo Repository, redis *redis.Client) Service {
	return NewService(repo, redis)
}

func NewServiceFactory() ServiceFactory {
	return &BaseServiceFactory{}
}
