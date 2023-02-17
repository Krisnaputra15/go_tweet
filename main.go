package main

import (
	"github.com/Krisnaputra15/go-tweet-echo/config"
	"github.com/Krisnaputra15/go-tweet-echo/db"
	"github.com/Krisnaputra15/go-tweet-echo/route"
	"github.com/labstack/echo/v4"
)

func init() {
	config.LoadEnv()
	db.ConnectToDB()
}

func main() {
	e := echo.New()

	route.InitRoute()

	e.Logger.Fatal(e.Start(":3000"))
}