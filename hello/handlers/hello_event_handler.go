package handlers

import (
	"context"

	"github.com/y16ra/deploy-lambda-go-container/hello/events"
)

func HandleHelloEvent(ctx context.Context, event *events.HelloEvent) (string, error) {
	if event == nil || len(event.Name) == 0 {
		return "Hello, World!", nil
	}
	return "Hello, " + event.Name + "!", nil
}
