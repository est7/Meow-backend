package models

import (
	"github.com/google/uuid"
	"time"
)

//	BaseEntity
//
// ID: 使用 UUID 作为主键，这是一种常见的做法，特别是在分布式系统中。
// CreatedAt: 记录实体创建的时间。
// UpdatedAt: 记录实体最后更新的时间。
// DeletedAt: 用于软删除，记录删除时间。使用指针类型允许为 nil。
// IsLogicDeleted: 布尔值，表示是否逻辑删除。
// Version: 整数类型，用于乐观锁或版本控制。
type BaseEntity struct {
	ID             uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt      time.Time  `json:"created_at" gorm:"type:timestamp with time zone;default:current_timestamp"`
	UpdatedAt      time.Time  `json:"updated_at" gorm:"type:timestamp with time zone;default:current_timestamp"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	IsLogicDeleted bool       `json:"is_logic_deleted" gorm:"default:false"`
	Version        int        `json:"version" gorm:"default:1"`
}
