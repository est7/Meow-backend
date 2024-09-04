package interfaces

import "Meow-backend/internal/models"

type CommonUserService interface {
	CheckUsernameExists(username string) (bool, error)
	CheckPhoneIsExist(phone string) (bool, error)
	SendEmailVerificationCode(email string) error
	SendPhoneVerificationCode(phone string) error
	VerifyEmailCode(email, code string) error
	VerifyPhoneCode(phone, code string) error

	GetUserByID(userID int64) (*models.User, error)
}
