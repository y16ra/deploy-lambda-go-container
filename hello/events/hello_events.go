package events

type HelloEvent struct {
	Name string `json:"name"`
}

func (e *HelloEvent) GetName() string {
	if len(e.Name) == 0 {
		return "World"
	}
	return e.Name
}
