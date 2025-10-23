package event

type Event struct {
	Type string
	Data any
}

type EventBus struct {
	bus chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		bus: make(chan Event),
	}
}

func (eb *EventBus) Publish(e Event) {
	eb.bus <- e
}

func (eb *EventBus) Subscribe() <-chan Event {
	return eb.bus
}
