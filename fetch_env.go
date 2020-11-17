package env

import (
	"github.com/joho/godotenv"
	"log"
)

var myEnv map[string]string
var err error
var isFetched bool = false

func initializeEnv() (map[string]string, error) {
	var envVariables map[string]string
	var err error

	if !isFetched {
		log.Println("Fetching variables from .env file for the first time.")
		envVariables, err = godotenv.Read()

		if err != nil {
			return nil, err
		}

		isFetched = true
	}

	return envVariables, err
}
