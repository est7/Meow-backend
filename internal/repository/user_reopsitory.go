package repository

import (
	"Meow-backend/internal/interfaces"
	"Meow-backend/internal/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.UserBaseModel, error)
	CreateUser(user *models.UserBaseModel) error
	GetUserByID(id int) (*models.UserBaseModel, error)
	UpdateUser(user *models.UserBaseModel) error
	DeleteUser(user *models.UserBaseModel) error
}

type UserRepositoryImpl struct {
	interfaces.Repository
}

func NewUserRepository(base interfaces.Repository) *UserRepositoryImpl {
	return &UserRepositoryImpl{Repository: base}
}

func (repo *UserRepositoryImpl) GetUserByEmail(email string) (*models.UserBaseModel, error) {
	var user models.UserBaseModel
	if err := repo.GetDB().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (repo *UserRepositoryImpl) CreateUser(user *models.UserBaseModel) error {
	return repo.GetDB().Create(user).Error
}
func (repo *UserRepositoryImpl) GetUserByID(id int) (*models.UserBaseModel, error) {
	var user models.UserBaseModel
	if err := repo.GetDB().Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (repo *UserRepositoryImpl) UpdateUser(user *models.UserBaseModel) error {
	return repo.GetDB().Save(user).Error
}
func (repo *UserRepositoryImpl) DeleteUser(user *models.UserBaseModel) error {
	return repo.GetDB().Delete(user).Error
}
