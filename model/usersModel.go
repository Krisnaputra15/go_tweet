package model

type User struct {
	Base
	Email          string     `gorm:"unique" json:"email"`
	Username       string     `gorm:"unique" json:"username"`
	Password       string     `json:"password"`
	Bio            string     `json:"bio"`
	Location       string     `json:"location"`
	Website        string     `json:"website"`
	FollowersCount int32      `json:"followers_count"`
	FollowingCount int32      `json:"following_count"`
	Tweets         []Tweet    `gorm:"foreignKey:UserID;references:ID;OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Followings     []Follower `gorm:"foreignKey:FollowingID;references:ID;OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FollowFrom     Follower   `gorm:"foreignKey:UserID;references:ID;OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
