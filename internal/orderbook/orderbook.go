package orderbook

import (
	"sort"
	"sync"

	"gox_engine/internal/order"
)

type OrderBook struct {
	mu sync.RWMutex

	Bids map[float64][]*order.Order
	Asks map[float64][]*order.Order

	BidPrices []float64
	AskPrices []float64

	Orders map[string]*order.Order
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		Bids:   make(map[float64][]*order.Order),
		Asks:   make(map[float64][]*order.Order),
		Orders: make(map[string]*order.Order),
	}
}

// Safe wrapper for matching logic
func (ob *OrderBook) ExecuteMatch(fn func()) {
	ob.mu.Lock()
	defer ob.mu.Unlock()
	fn()
}

func (ob *OrderBook) AddOrder(o *order.Order) {
	ob.mu.Lock()
	defer ob.mu.Unlock()

	ob.Orders[o.ID] = o

	if o.Side == order.Buy {
		ob.Bids[o.Price] = append(ob.Bids[o.Price], o)
		ob.UpdateBidPrices()
	} else {
		ob.Asks[o.Price] = append(ob.Asks[o.Price], o)
		ob.UpdateAskPrices()
	}
}

func (ob *OrderBook) CancelOrder(orderID string) bool {
	ob.mu.Lock()
	defer ob.mu.Unlock()

	o, exists := ob.Orders[orderID]
	if !exists {
		return false
	}

	var book map[float64][]*order.Order

	if o.Side == order.Buy {
		book = ob.Bids
	} else {
		book = ob.Asks
	}

	queue := book[o.Price]

	for i, ord := range queue {
		if ord.ID == orderID {
			book[o.Price] = append(queue[:i], queue[i+1:]...)
			delete(ob.Orders, orderID)
			return true
		}
	}

	return false
}

func (ob *OrderBook) UpdateBidPrices() {
	ob.BidPrices = ob.BidPrices[:0]
	for price := range ob.Bids {
		ob.BidPrices = append(ob.BidPrices, price)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(ob.BidPrices)))
}

func (ob *OrderBook) UpdateAskPrices() {
	ob.AskPrices = ob.AskPrices[:0]
	for price := range ob.Asks {
		ob.AskPrices = append(ob.AskPrices, price)
	}
	sort.Float64s(ob.AskPrices)
}
