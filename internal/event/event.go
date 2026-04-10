package event

import "time"

type EventType string

const (
	OrderPlaced   EventType = "OrderPlaced"
	OrderMatched  EventType = "OrderMatched"
	OrderCanceled EventType = "OrderCanceled"
)

type Event struct {
	Type      EventType
	Data      interface{}
	Timestamp time.Time
}
