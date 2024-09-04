package servicefactory

import (
	"Meow-backend/internal/interfaces"
)

type IMServiceFactory struct {
	interfaces.ServiceFactory
}

func NewIMServiceFactory(base interfaces.ServiceFactory) *IMServiceFactory {
	return &IMServiceFactory{ServiceFactory: base}
}

//func (f *IMServiceFactory) CreateService(repo interfaces.Repository, redis *redis.Client) interfaces.Service {
//	baseService := f.ServiceFactory.CreateService(repo, redis)
//	return service.NewIMService(baseService, f.CreateCommonUserService(repo, redis))
//}

//func (f *IMServiceFactory) CreateCommonUserService(repo interfaces.Repository, client *redis.Client) interfaces.CommonUserService {
//	return NewCommonUserService(repo, client)
//}
