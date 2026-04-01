package orderbook

import (
	"gox_engine/internal/order"
	"sort"
)

type OrderBook struct {
Bids map[float64][]*order.Order
Asks map[float64][]*order.Order

BidPrices []float64
AskPrices []float64
}

func NewOrderBook() *OrderBook {
return &OrderBook{
Bids: make(map[float64][]*order.Order),
Asks: make(map[float64][]*order.Order),
}
}

func (ob *OrderBook) AddOrder(o *order.Order) {
if o.Side == order.Buy {
ob.Bids[o.Price] = append(ob.Bids[o.Price], o)
ob.UpdateBidPrices()
} else {
ob.Asks[o.Price] = append(ob.Asks[o.Price], o)
ob.UpdateAskPrices()
}
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
