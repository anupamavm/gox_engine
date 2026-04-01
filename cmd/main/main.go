package main

import (
"fmt"
"time"

"gox_engine/internal/matching"
"gox_engine/internal/order"
"gox_engine/internal/orderbook"
)

func main() {
ob := orderbook.NewOrderBook()

// Add sell orders
s1 := &order.Order{
ID:        "S1",
Type:      order.Limit,
Side:      order.Sell,
Price:     100,
Quantity:  10,
Timestamp: time.Now(),
}

s2 := &order.Order{
ID:        "S2",
Type:      order.Limit,
Side:      order.Sell,
Price:     101,
Quantity:  5,
Timestamp: time.Now(),
}

ob.AddOrder(s1)
ob.AddOrder(s2)

// Incoming buy order
buy := &order.Order{
ID:        "B1",
Type:      order.Market,
Side:      order.Buy,
Quantity:  12,
Timestamp: time.Now(),
}

trades := matching.Match(ob, buy)

for _, t := range trades {
fmt.Printf("TRADE: %+v\n", t)
}
}
