package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type HelloEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event *HelloEvent) (string, error) {
	if event == nil {
		return "Hello, World!", nil
	}
	return "Hello, " + event.Name + "!", nil
}

func main() {
	lambda.Start(HandleRequest)
}
