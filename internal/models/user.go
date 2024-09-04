package models

import "time"

type User struct {
	BaseEntity
	Username          string    `json:"username" gorm:"type:varchar(50);uniqueIndex;not null"`
	Email             string    `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Password          string    `json:"-" gorm:"type:varchar(255);not null"` // 不输出到JSON
	FirstName         string    `json:"first_name" gorm:"type:varchar(50)"`
	LastName          string    `json:"last_name" gorm:"type:varchar(50)"`
	Avatar            string    `json:"avatar" gorm:"type:varchar(255)"`
	Bio               string    `json:"bio" gorm:"type:text"`
	DateOfBirth       time.Time `json:"date_of_birth" gorm:"type:date"`
	PhoneNumber       string    `json:"phone_number" gorm:"type:varchar(20)"`
	LastLoginAt       time.Time `json:"last_login_at" gorm:"type:timestamp with time zone"`
	IsVerified        bool      `json:"is_verified" gorm:"default:false"`
	Role              string    `json:"role" gorm:"type:varchar(20);default:'user'"`
	PreferredLanguage string    `json:"preferred_language" gorm:"type:varchar(10);default:'en'"`
}
