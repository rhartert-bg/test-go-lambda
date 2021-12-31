package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/rhartert-bg/test-go-lambda/internal/lib"
)

type MyEvent struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	return lib.FormatMessage(event.Message), nil
}

func main() {
	lambda.Start(HandleRequest)
}
