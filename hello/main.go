package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/y16ra/deploy-lambda-go-container/hello/handlers"
)

func main() {
	lambda.Start(handlers.HandleHelloEvent)
}
