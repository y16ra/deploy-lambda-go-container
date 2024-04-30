package handlers

import (
	"context"
	"testing"

	"github.com/y16ra/deploy-lambda-go-container/hello/events"
)

func TestHandleHelloEvent(t *testing.T) {
	type args struct {
		event *events.HelloEvent
	}
	tests := map[string]struct {
		args    args
		want    string
		wantErr bool
	}{
		"normal": {
			args:    args{event: &events.HelloEvent{Name: "Alice"}},
			want:    "Hello, Alice!",
			wantErr: false,
		},
		"empty": {
			args:    args{event: &events.HelloEvent{Name: ""}},
			want:    "Hello, World!",
			wantErr: false,
		},
		"nil": {
			args:    args{event: nil},
			want:    "Hello, World!",
			wantErr: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			got, err := HandleHelloEvent(ctx, tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleHelloEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HandleHelloEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
