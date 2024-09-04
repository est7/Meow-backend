package service

import (
	"Meow-backend/internal/interfaces"
	"Meow-backend/internal/models"
	"Meow-backend/internal/modules/v1/im/repository"
)

type IMService interface {
	interfaces.CommonUserService
	PrivateIMService
}

type IMServiceImpl struct {
	interfaces.Service
	CommonUserService interfaces.CommonUserService
	imRepo            *repository.IMRepositoryImpl
}

type PrivateIMService interface {
	EnterChatRoom()
}

func NewIMService(base interfaces.Service, commonUserService interfaces.CommonUserService) *IMServiceImpl {
	return &IMServiceImpl{
		Service:           base,
		CommonUserService: commonUserService,
		imRepo:            repository.NewIMRepository(base.GetRepo()),
	}
}

func (s *IMServiceImpl) EnterChatRoom() {
	err := s.imRepo.CreateChatRoom(&models.BaseEntity{})
	if err != nil {
		return
	}

}
