package servicefactory

import (
	"Meow-backend/internal/interfaces"
	"Meow-backend/internal/modules/v1/user/service"
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
	return service.NewUserService(baseService)
}

func (f *UserServiceFactory) CreateCommonUserService(repo interfaces.Repository, redis *redis.Client) interfaces.CommonUserService {
	baseService := f.ServiceFactory.CreateService(repo, redis)
	return service.NewCommonUserService(baseService)
}
