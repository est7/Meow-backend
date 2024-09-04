package models

import (
	"github.com/google/uuid"
	"time"
)

type Card struct {
	BaseEntity
	UserID         uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	CardNumber     string    `json:"card_number" gorm:"type:varchar(20);uniqueIndex;not null"`
	CardholderName string    `json:"cardholder_name" gorm:"type:varchar(100);not null"`
	ExpirationDate time.Time `json:"expiration_date" gorm:"type:date;not null"`
	CVV            string    `json:"-" gorm:"type:varchar(4);not null"` // 不输出到JSON
	CardType       string    `json:"card_type" gorm:"type:varchar(20)"`
	BillingAddress string    `json:"billing_address" gorm:"type:text"`
	IsDefault      bool      `json:"is_default" gorm:"default:false"`
	LastFourDigits string    `json:"last_four_digits" gorm:"type:varchar(4)"`
	User           User      `json:"-" gorm:"foreignKey:UserID"`
}
