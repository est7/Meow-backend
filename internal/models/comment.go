package models

import "github.com/google/uuid"

type Comment struct {
	BaseEntity
	FeedID          uuid.UUID  `json:"feed_id" gorm:"type:uuid;not null"`
	UserID          uuid.UUID  `json:"user_id" gorm:"type:uuid;not null"`
	Content         string     `json:"content" gorm:"type:text;not null"`
	LikeCount       int        `json:"like_count" gorm:"default:0"`
	ParentCommentID *uuid.UUID `json:"parent_comment_id" gorm:"type:uuid"`
	IsEdited        bool       `json:"is_edited" gorm:"default:false"`
	Feed            Feed       `json:"-" gorm:"foreignKey:FeedID"`
	User            User       `json:"-" gorm:"foreignKey:UserID"`
}
