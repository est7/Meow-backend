package service

import (
	"Meow-backend/internal/models"
	"Meow-backend/pkg/crypto"
	"Meow-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (s *UserServiceImpl) RegisterWithEmail(c *gin.Context, email string, emailVCode string) error {
	// 验证邮箱验证码
	if err := s.VerifyEmailCode(email, emailVCode); err != nil {
		return err
	}

	// 检查邮箱是否已被绑定
	isExisted, err := s.userRepo.CheckEmailIsExist(email)
	if err != nil {
		return err
	}
	if isExisted {
		return errors.New("email already exists")
	}

	// 创建用户
	user := &models.User{
		Email: email,
		// 生成随机用户名，或者让用户在后续步骤中设置
		Username: utils.GenerateRandomUsername(),
		// 这里应该有一个安全的密码生成方法，或者让用户在后续步骤中设置密码
		Password: utils.GenerateSecurePassword(),
	}

	return s.userRepo.CreateUser(user)
}

func (s *UserServiceImpl) RegisterWithUsername(c *gin.Context, username string, password string) error {
	// 检查用户名是否已被使用
	isExisted, err := s.userRepo.CheckUsernameExists(username)
	if err != nil {
		return err
	}
	if isExisted {
		return errors.New("username already exists")
	}

	// 创建用户
	user := &models.User{
		Username: username,
		Password: crypto.HashPassword(password), // 使用适当的密码哈希函数
	}

	return s.userRepo.CreateUser(user)
}

func (s *UserServiceImpl) RegisterWithPhone(c *gin.Context, phone string, phoneVCode string) error {
	// 验证手机验证码
	if err := s.verifyPhoneCode(phone, phoneVCode); err != nil {
		return err
	}

	// 检查手机号是否已被绑定
	isExisted, err := s.userRepo.CheckPhoneIsExist(phone)
	if err != nil {
		return err
	}
	if isExisted {
		return errors.New("phone number already exists")
	}

	// 创建用户
	user := &models.User{
		PhoneNumber: phone,
		// 生成随机用户名，或者让用户在后续步骤中设置
		Username: utils.GenerateRandomUsername(),
		// 这里应该有一个安全的密码生成方法，或者让用户在后续步骤中设置密码
		Password: utils.GenerateSecurePassword(),
	}

	return s.userRepo.CreateUser(user)
}
