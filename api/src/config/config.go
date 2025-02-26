package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)
var(
	Port int
	StringConnectionDataBase string
)

func Load() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 8080
	}

	StringConnectionDataBase = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)


}
