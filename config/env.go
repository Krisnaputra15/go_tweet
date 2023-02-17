package config

import (
	"log"

	"github.com/joho/godotenv"
)

var envkey map[string]string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
}
