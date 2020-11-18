package env

import (
	"github.com/joho/godotenv"
	"log"
)

var isFetched bool = false

func initializeEnv() error {
	if !isFetched {
		err := godotenv.Load()
		log.Println("Fetching variables from .env file for the first time.")

		if err != nil {
			return err
		}

		isFetched = true
	}

	return nil
}
