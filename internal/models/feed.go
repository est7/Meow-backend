package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Feed struct {
	BaseEntity
	UserID       uuid.UUID      `json:"user_id" gorm:"type:uuid;not null"`
	Content      string         `json:"content" gorm:"type:text;not null"`
	MediaURL     string         `json:"media_url" gorm:"type:varchar(255)"`
	MediaType    string         `json:"media_type" gorm:"type:varchar(20)"`
	LikeCount    int            `json:"like_count" gorm:"default:0"`
	CommentCount int            `json:"comment_count" gorm:"default:0"`
	ShareCount   int            `json:"share_count" gorm:"default:0"`
	IsPublic     bool           `json:"is_public" gorm:"default:true"`
	Location     string         `json:"location" gorm:"type:varchar(100)"`
	Tags         pq.StringArray `json:"tags" gorm:"type:text[]"`
	Topics       pq.StringArray `json:"topics" gorm:"type:text[]"`
	User         User           `json:"-" gorm:"foreignKey:UserID"`
}
