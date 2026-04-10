package event

import "sync"

type EventStore struct {
	mu     sync.Mutex
	Events []Event
}

func NewEventStore() *EventStore {
	return &EventStore{
		Events: make([]Event, 0),
	}
}

func (es *EventStore) Append(e Event) {
	es.mu.Lock()
	defer es.mu.Unlock()

	es.Events = append(es.Events, e)
}

func (es *EventStore) GetAll() []Event {
	es.mu.Lock()
	defer es.mu.Unlock()

	return append([]Event(nil), es.Events...)
}
