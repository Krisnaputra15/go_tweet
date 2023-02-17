package model

type Tweet struct {
	Base
	Content      string    `json:"content"`
	ViewsCount   int32     `json:"views"`
	LikesCount   int32     `json:"likes_count"`
	RetweetCount int32     `json:"retweet_count"`
	RepliedTo    string    `gorm:"size:191" json:"replied_to"`
	UserID       string    `gorm:"size:191" json:"user_id"`
	Replies      []Tweet   `gorm:"foreignKey:RepliedTo;references:ID;OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Likes        []Like    `gorm:"foreignKey:TweetID;references:ID;OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Retweets     []Retweet `gorm:"foreignKey:TweetID;references:ID;OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Medias       []Media   `gorm:"foreignKey:TweetID;references:ID"`
	User         User      `gorm:"foreignKey:UserID"`
}
