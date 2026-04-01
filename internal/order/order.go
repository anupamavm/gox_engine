package order

import "time"

type Side string

const (
Buy  Side = "BUY"
Sell Side = "SELL"
)

type OrderType string

const (
Limit  OrderType = "LIMIT"
Market OrderType = "MARKET"
)

type Order struct {
ID        string
Type      OrderType
Side      Side
Price     float64
Quantity  int
Timestamp time.Time
}
