package env

import (
	"log"
	"os"
)

// GetEnvVar get the environment variables for the user
func GetEnvVar(key string, options Options) string {
	if !isFetched {
		err := InitializeEnv()

		if err != nil {
			log.Fatalf("An error occurred loading the .env file. \n%s\n", err)
		}
	}

	environment := os.Getenv("ENVIRONMENT")

	var fallback string

	if environment == "development" || environment == "" {
		fallback = options.DevDefault
	}

	value := os.Getenv(key)

	if value == "" {
		value = fallback
	}

	if !options.Optional && value == "" {
		log.Fatalf("%s is not defined in .env", key)
	}

	return value
}
