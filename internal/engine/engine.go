package engine

import (
	"time"

	"gox_engine/internal/event"
	"gox_engine/internal/matching"
	"gox_engine/internal/order"
	"gox_engine/internal/orderbook"
)

type Engine struct {
	OrderBook  *orderbook.OrderBook
	EventStore *event.EventStore

	OrderChan  chan *order.Order
	CancelChan chan string
}

func NewEngine() *Engine {
	return &Engine{
		OrderBook:  orderbook.NewOrderBook(),
		EventStore: event.NewEventStore(),
		OrderChan:  make(chan *order.Order, 100),
		CancelChan: make(chan string, 100),
	}
}

func (e *Engine) Start() {
	go func() {
		for {
			select {

			// ORDER FLOW
			case o := <-e.OrderChan:

				// Record event
				e.EventStore.Append(event.Event{
					Type:      event.OrderPlaced,
					Data:      o,
					Timestamp: time.Now(),
				})

				// Match
				trades := matching.Match(e.OrderBook, o)

				// Record trades
				for _, t := range trades {
					e.EventStore.Append(event.Event{
						Type:      event.OrderMatched,
						Data:      t,
						Timestamp: time.Now(),
					})
				}

			// CANCEL FLOW
			case id := <-e.CancelChan:

				success := e.OrderBook.CancelOrder(id)

				if success {
					e.EventStore.Append(event.Event{
						Type:      event.OrderCanceled,
						Data:      id,
						Timestamp: time.Now(),
					})
				}
			}
		}
	}()
}
