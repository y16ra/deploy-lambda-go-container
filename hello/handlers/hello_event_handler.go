package handlers

import (
	"context"
	"errors"

	"github.com/y16ra/deploy-lambda-go-container/hello/events"
)

func HandleHelloEvent(ctx context.Context, event *events.HelloEvent) (string, error) {
	if event == nil {
		return "", errors.New("event is nil")
	}
	return "Hello, " + event.GetName() + "!", nil
}
