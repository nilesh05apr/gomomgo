package configs

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading enviroment variables")
	}
	return os.Getenv("MONGOURI")
}