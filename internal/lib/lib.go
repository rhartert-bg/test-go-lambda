package lib

import (
	"fmt"
	"os"
)

func getEnv() string {
	if e := os.Getenv("DEPLOY_ENV"); e != "" {
		return e
	}
	return "UNKNOWN"
}

func FormatMessage(message string) string {
	return fmt.Sprintf("[Env: %s] Message: %s", getEnv(), message)
}
