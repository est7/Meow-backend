package repository

import (
	"Meow-backend/internal/interfaces"
	"Meow-backend/internal/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

type UserRepositoryImpl struct {
	interfaces.Repository
}

func NewUserRepository(base interfaces.Repository) *UserRepositoryImpl {
	return &UserRepositoryImpl{Repository: base}
}

// CreateUser 创建用户
func (repo *UserRepositoryImpl) CreateUser(user *models.User) error {
	return repo.GetDB().Create(user).Error

}

// GetUserByEmail 根据邮箱获取用户
func (repo *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.GetDB().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CheckUsernameExists 检查用户名是否存在
func (repo *UserRepositoryImpl) CheckUsernameExists(username string) (bool, error) {
	var count int64
	err := repo.GetDB().Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

// CheckPhoneIsExist 检查手机号是否存在
func (repo *UserRepositoryImpl) CheckPhoneIsExist(phone string) (bool, error) {
	var count int64
	err := repo.GetDB().Model(&models.User{}).Where("phone_number = ?", phone).Count(&count).Error
	return count > 0, err
}

// CheckEmailIsExist 检查邮箱是否存在
func (repo *UserRepositoryImpl) CheckEmailIsExist(email string) (bool, error) {
	var user models.User
	if err := repo.GetDB().Where("email = ?", email).First(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}

// GetUserByID 根据ID获取用户
func (repo *UserRepositoryImpl) GetUserByID(id int) (*models.User, error) {
	var user models.User
	if err := repo.GetDB().Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser 更新用户
func (repo *UserRepositoryImpl) UpdateUser(user *models.User) error {
	return repo.GetDB().Save(user).Error
}

// DeleteUser 删除用户
func (repo *UserRepositoryImpl) DeleteUser(user *models.User) error {
	return repo.GetDB().Delete(user).Error
}
