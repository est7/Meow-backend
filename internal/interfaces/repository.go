package interfaces

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetDB() *gorm.DB
}

type BaseRepository struct {
	DB *gorm.DB
}

func (r *BaseRepository) GetDB() *gorm.DB {
	return r.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &BaseRepository{DB: db}
}
