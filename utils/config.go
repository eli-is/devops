// B"H

package utils

// credit https://blog.risingstack.com/golang-tutorial-for-nodejs-developers-getting-started/#nethttp

import (
	"fmt"
	"go-http-server/structs"
	"os"
	"strconv"
)

func ReadConfig() structs.Config {

	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8000"
	}

	debugMode := os.Getenv("DEBUG_MODE")
	if debugMode == "" {
		debugMode = "true"
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "local"
	}

	port, err := strconv.Atoi(portString)

	if err != nil {
		panic(fmt.Sprintf("Could not parse %s to int", portString))
	}

	return structs.Config{
		Port:        port,
		DebugMode:   debugMode,
		Environment: environment,
	}
}
