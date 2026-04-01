package orderbook

import (
	"gox/internal/order"
)

type OrderBook struct {
	Bids map[float64][]*order.Order
	Asks map[float64][]*order.Order

	BidPrices []float64
	AskPrices []float64
}