package factoryforservice

import (
	"Meow-backend/internal/interfaces"
	"Meow-backend/internal/modules/v1/user/service"
	"github.com/redis/go-redis/v9"
)

type UserServiceFactory struct {
}

func NewUserServiceFactory() *UserServiceFactory {
	return &UserServiceFactory{}
}

func (f *UserServiceFactory) CreateService(repo interfaces.Repository, redis *redis.Client) service.UserService {
	baseService := interfaces.NewBaseService(repo, redis)
	return service.NewUserService(baseService)
}
