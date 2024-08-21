package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username     string    `json:"username" gorm:"type:varchar(255);not null;unique"`
	Email        string    `json:"email" gorm:"type:varchar(255);unique"`
	PasswordHash string    `json:"-" gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp with time zone;default:current_timestamp"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp with time zone;default:current_timestamp"`
}
