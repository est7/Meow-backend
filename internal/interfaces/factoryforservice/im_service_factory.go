package factoryforservice

import (
	"Meow-backend/internal/interfaces"
	imService "Meow-backend/internal/modules/v1/im/service"
	userService "Meow-backend/internal/modules/v1/user/service"
	"github.com/redis/go-redis/v9"
)

type IMServiceFactory struct{}

func NewIMServiceFactory() *IMServiceFactory {
	return &IMServiceFactory{}
}

func (f *IMServiceFactory) CreateService(repo interfaces.Repository, redis *redis.Client) imService.IMService {
	baseService := interfaces.NewBaseService(repo, redis)
	commonUserService := userService.NewCommonUserService(repo, redis)
	return imService.NewIMService(baseService, commonUserService, repo)
}
