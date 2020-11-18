package env

import (
	"github.com/joho/godotenv"
	"log"
)

var isFetched bool = false

// InitializeEnv method for initializing .env
func InitializeEnv() error {
	if !isFetched {
		log.Println("Fetching variables from .env file for the first time.")
		err := godotenv.Load()

		isFetched = true

		if err != nil {
			log.Printf("%s", err)
			return err
		}
	}

	return nil
}
