package model

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	TweetID string `gorm:"size:191" json:"tweet_id"`
	UserID  string `gorm:"size:191" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID"`
	Tweet   Tweet  `gorm:"foreignKey:TweetID"`
}
