package events

import "testing"

func TestHelloEvent_GetName(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := map[string]struct {
		fields fields
		want   string
	}{
		"normal": {
			fields: fields{Name: "Alice"},
			want:   "Alice",
		},
		"empty": {
			fields: fields{Name: ""},
			want:   "World",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			e := &HelloEvent{
				Name: tt.fields.Name,
			}
			if got := e.GetName(); got != tt.want {
				t.Errorf("HelloEvent.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
