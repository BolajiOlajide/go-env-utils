package env

import (
	"github.com/joho/godotenv"
	"log"
)

var isFetched bool = false

func initializeEnv() error {
	if !isFetched {
		log.Println("Fetching variables from .env file for the first time.")
		err := godotenv.Load()

		if err != nil {
			return err
		}

		isFetched = true
	}

	return nil
}
