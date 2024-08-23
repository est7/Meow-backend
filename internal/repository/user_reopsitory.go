package repository

import (
	"Meow-backend/internal/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.UserBaseModel, error)
	CreateUser(user *models.UserBaseModel) error
	GetUserByID(id int) (*models.UserBaseModel, error)
	UpdateUser(user *models.UserBaseModel) error
	DeleteUser(user *models.UserBaseModel) error
}

type UserRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewUserRepositoryImpl(db *gorm.DB, redis *redis.Client) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

// GetUserByEmail 实现
func (r *UserRepositoryImpl) GetUserByEmail(email string) (*models.UserBaseModel, error) {
	var user *models.UserBaseModel
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.UserBaseModel{}, nil
}

// CreateUser 实现
func (r *UserRepositoryImpl) CreateUser(user *models.UserBaseModel) error {
	return r.db.Create(user).Error
}

// GetUserByID 实现
func (r *UserRepositoryImpl) GetUserByID(id int) (*models.UserBaseModel, error) {
	var user *models.UserBaseModel
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.UserBaseModel{}, nil
}

// UpdateUser 实现
func (r *UserRepositoryImpl) UpdateUser(user *models.UserBaseModel) error {
	return r.db.Save(user).Error
}

// DeleteUser 实现
func (r *UserRepositoryImpl) DeleteUser(user *models.UserBaseModel) error {
	return r.db.Delete(user).Error
}
