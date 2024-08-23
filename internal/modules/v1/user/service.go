package user

import (
	"Meow-backend/internal/repository"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
	rc *redis.Client
}

func newService(db *gorm.DB, rc *redis.Client) *Service {
	return &Service{
		db: db,
		rc: rc,
	}
}

type UserService struct {
	*Service // 嵌入 Service
	repo     repository.UserRepository
}

func NewUserService(db *gorm.DB, rc *redis.Client, repo repository.UserRepository) *UserService {
	return &UserService{
		Service: newService(db, rc),
		repo:    repo,
	}
}
