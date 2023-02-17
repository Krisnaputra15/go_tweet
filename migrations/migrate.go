package main

import (
	"github.com/Krisnaputra15/go-tweet-echo/config"
	"github.com/Krisnaputra15/go-tweet-echo/db"
	"github.com/Krisnaputra15/go-tweet-echo/model"
)

func init() {
	config.LoadEnv()
	db.ConnectToDB()
}

func main() {
	db.DB.AutoMigrate(&model.User{}, &model.Tweet{}, &model.Like{}, &model.Retweet{}, &model.Media{}, &model.Follower{})
}