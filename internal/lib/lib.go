package lib

import (
	"fmt"
	"os"
	"strings"
)

func getEnv(envVar string) string {
	if e := os.Getenv(envVar); e != "" {
		return e
	}
	return "UNKNOWN"
}

func FormatMessage(message string) string {
	sb := strings.Builder{}
	sb.WriteString("New!\n")
	sb.WriteString(fmt.Sprintf("Message: %s\n", message))
	sb.WriteString(fmt.Sprintf("Environment: %s\n", getEnv("BG_DEPLOYMENT_ENVIRONMENT")))
	sb.WriteString(fmt.Sprintf("Git Hash: %s\n", getEnv("BG_BUILD_GIT_HASH")))
	return sb.String()
}
