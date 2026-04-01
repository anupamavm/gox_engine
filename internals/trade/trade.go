package trade

import "time"

type Trade struct {
	ID         string
	BuyOrderID  string
	SellOrderID string
	Price      float64
	Quantity   int
	Timestamp  time.Time
}