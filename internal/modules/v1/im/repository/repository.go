package repository

import (
	"Meow-backend/internal/interfaces"
	"Meow-backend/internal/models"
)

type IMRepository interface {
	CreateChatRoom(room *models.BaseEntity) error
}

type IMRepositoryImpl struct {
	interfaces.Repository
}

func NewIMRepository(base interfaces.Repository) *IMRepositoryImpl {
	return &IMRepositoryImpl{Repository: base}
}

func (repo *IMRepositoryImpl) CreateChatRoom(IM *models.BaseEntity) error {
	return repo.GetDB().Create(IM).Error

}
