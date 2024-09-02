package service

import (
	"Meow-backend/internal/interfaces"
	"github.com/redis/go-redis/v9"
)

type UserServiceFactory struct {
	interfaces.ServiceFactory
}

func NewUserServiceFactory(base interfaces.ServiceFactory) *UserServiceFactory {
	return &UserServiceFactory{ServiceFactory: base}
}

func (f *UserServiceFactory) CreateService(repo interfaces.Repository, redis *redis.Client) interfaces.Service {
	baseService := f.ServiceFactory.CreateService(repo, redis)
	return NewUserService(baseService)
}
