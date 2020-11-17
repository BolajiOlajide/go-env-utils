package env

import (
	"fmt"
	"log"
)

// GetEnvVar get the environment variables for the user
func GetEnvVar(key string, options Options) (string, error) {
	envVariables, err := initializeEnv()

	if err != nil {
		log.Fatalf("An error occurred loading the .env file. \n%s\n", err)
	}

	environment := envVariables["ENVIRONMENT"]

	var fallback string

	if environment == "development" || environment == "" {
		fallback = options.DevDefault
	}

	value := envVariables[key]

	if value == "" {
		value = fallback
	}

	if !options.Optional && value == "" {
		return "", fmt.Errorf("%s is not defined in .env", key)
	}

	return value, nil
}
