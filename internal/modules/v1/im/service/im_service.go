package service

import (
	"Meow-backend/internal/interfaces"
	"Meow-backend/internal/models"
	"Meow-backend/internal/modules/v1/im/repository"
)

type IMService interface {
	interfaces.BaseService
	interfaces.CommonUserService
	PrivateIMService
}

type IMServiceImpl struct {
	interfaces.BaseService
	interfaces.CommonUserService
	imRepo *repository.IMRepositoryImpl
}

func NewIMService(base interfaces.BaseService, commonUserService interfaces.CommonUserService, repo interfaces.Repository) IMService {
	return &IMServiceImpl{
		BaseService:       base,
		CommonUserService: commonUserService,
		imRepo:            repository.NewIMRepository(repo),
	}
}

type PrivateIMService interface {
	EnterChatRoom() error
}

func (s *IMServiceImpl) EnterChatRoom() error {
	user, err := s.GetUserByID(313)
	if err != nil {
		return err
	}

	// Use the user object as needed
	return s.imRepo.CreateChatRoom(&models.BaseEntity{})
}
