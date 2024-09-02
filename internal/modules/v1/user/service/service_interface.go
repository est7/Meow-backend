package service

import (
	"Meow-backend/internal/interfaces"
	"Meow-backend/internal/models"
	"Meow-backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type IUserService interface {
	RegisterUser(username, email, password string) error
	CheckUsernameExists(username string) (bool, error)
	GetUserByID(userID int64) (*models.UserBaseModel, error)
	RegisterWithEmail(c *gin.Context, email string, password string) error
	RegisterWithUsername(c *gin.Context, username string, password string) error
}

type UserServiceImpl struct {
	interfaces.Service
	userRepo *repository.UserRepositoryImpl
}

func NewUserService(base interfaces.Service) *UserServiceImpl {
	return &UserServiceImpl{
		Service:  base,
		userRepo: repository.NewUserRepository(base.GetRepo()),
	}
}

// RegisterUser 注册用户
func (s *UserServiceImpl) RegisterUser(username, email, password string) error {
	var user models.UserBaseModel
	user.Username = username
	user.Email = email
	user.PasswordHash = password
	return s.userRepo.CreateUser(&user)
}

func (s *UserServiceImpl) CheckUsernameExists(username string) (bool, error) {
	return false, nil
}

func (s *UserServiceImpl) GetUserByID(userID int64) (*models.UserBaseModel, error) {
	return s.userRepo.GetUserByID(int(userID))
}

func (s *UserServiceImpl) RegisterWithEmail(c *gin.Context, email string, password string) error {
	return nil
}

func (s *UserServiceImpl) RegisterWithUsername(c *gin.Context, username string, password string) error {
	return nil
}
