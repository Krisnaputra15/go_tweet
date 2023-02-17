package model

type Media struct {
	ID        int32  `json:"id"`
	MediaName string `json:"media_name"`
	TweetID   string `gorm:"size:191" json:"tweet_id"`
	Tweet     Tweet  `gorm:"foreignKey:TweetID"`
}
