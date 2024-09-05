package service

import (
	"Meow-backend/internal/interfaces"
	"Meow-backend/internal/models"
	"Meow-backend/internal/modules/v1/user/repository"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type PrivateUserService interface {
	RegisterWithUsername(ctx *gin.Context, username, password string) error
	RegisterWithEmail(ctx *gin.Context, email, emailVCode string) error
	RegisterWithPhone(ctx *gin.Context, phone, phoneVCode string) error
}

type UserService interface {
	interfaces.CommonUserService
	PrivateUserService
}

type UserServiceImpl struct {
	interfaces.BaseService
	userRepo *repository.UserRepositoryImpl
}

func NewUserService(base interfaces.BaseService) *UserServiceImpl {
	return &UserServiceImpl{
		BaseService: base,
		userRepo:    repository.NewUserRepository(base.GetRepo()),
	}
}

func NewCommonUserService(repo interfaces.Repository, client *redis.Client) interfaces.CommonUserService {
	return &CommonUserServiceImpl{
		repo:   repo,
		client: client,
	}
}

// 实现 CommonUserService 的方法

func (s *UserServiceImpl) SendEmailVerificationCode(email string) error {
	// 实现发送邮件验证码的逻辑
	//TODO implement me
	panic("implement me")
}

func (s *UserServiceImpl) SendPhoneVerificationCode(phone string) error {
	// 实现发送手机验证码的逻辑
	//TODO implement me
	panic("implement me")
}

func (s *UserServiceImpl) VerifyEmailCode(email, code string) error {
	// 实现验证邮箱验证码的逻辑
	//TODO implement me
	panic("implement me")
}

func (s *UserServiceImpl) VerifyPhoneCode(phone, code string) error {
	// 实现验证手机验证码的逻辑
	//TODO implement me
	panic("implement me")
}

func (s *UserServiceImpl) GetUserByID(userID int64) (*models.User, error) {
	return s.userRepo.GetUserByID(int(userID))
}

// RegisterUser 注册用户
func (s *UserServiceImpl) RegisterUser(username, email, password string) error {
	var user models.User
	user.Username = username
	user.Email = email
	user.Password = password
	return s.userRepo.CreateUser(&user)
}

func (s *UserServiceImpl) CheckUsernameExists(username string) (bool, error) {
	return false, nil
}

func (s *UserServiceImpl) CheckPhoneIsExist(phone string) (bool, error) {
	return false, nil
}
