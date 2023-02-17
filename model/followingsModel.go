package model

import "gorm.io/gorm"

type Follower struct {
	gorm.Model
	UserID      string `gorm:"size:191" json:"user_id"`
	FollowingID string `gorm:"size:191" json:"following_id"`
}
