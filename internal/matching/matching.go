package matching

import (
	"fmt"
	"time"

	"gox_engine/internal/order"
	"gox_engine/internal/orderbook"
	"gox_engine/internal/trade"
)

func Match(ob *orderbook.OrderBook, incoming *order.Order) []trade.Trade {
	var trades []trade.Trade

	if incoming.Side == order.Buy {
		trades = matchBuy(ob, incoming)
	} else {
		trades = matchSell(ob, incoming)
	}

	if incoming.Quantity > 0 && incoming.Type == order.Limit {
		ob.AddOrder(incoming)
	}

	return trades
}

func matchBuy(ob *orderbook.OrderBook, buy *order.Order) []trade.Trade {
	var trades []trade.Trade

	for len(ob.AskPrices) > 0 && buy.Quantity > 0 {
		bestAsk := ob.AskPrices[0]

		if buy.Type == order.Limit && buy.Price < bestAsk {
			break
		}

		queue := ob.Asks[bestAsk]

		for len(queue) > 0 && buy.Quantity > 0 {
			sell := queue[0]

			tradeQty := min(buy.Quantity, sell.Quantity)

			t := trade.Trade{
				ID:          fmt.Sprintf("%d", time.Now().UnixNano()),
				BuyOrderID:  buy.ID,
				SellOrderID: sell.ID,
				Price:       bestAsk,
				Quantity:    tradeQty,
				Timestamp:   time.Now(),
			}

			trades = append(trades, t)

			buy.Quantity -= tradeQty
			sell.Quantity -= tradeQty

			if sell.Quantity == 0 {
				queue = queue[1:]
			}
		}

		if len(queue) == 0 {
			delete(ob.Asks, bestAsk)
			ob.UpdateAskPrices()
		} else {
			ob.Asks[bestAsk] = queue
		}
	}

	return trades
}

func matchSell(ob *orderbook.OrderBook, sell *order.Order) []trade.Trade {
	var trades []trade.Trade

	for len(ob.BidPrices) > 0 && sell.Quantity > 0 {
		bestBid := ob.BidPrices[0]

		if sell.Type == order.Limit && sell.Price > bestBid {
			break
		}

		queue := ob.Bids[bestBid]

		for len(queue) > 0 && sell.Quantity > 0 {
			buy := queue[0]

			tradeQty := min(sell.Quantity, buy.Quantity)

			t := trade.Trade{
				ID:          fmt.Sprintf("%d", time.Now().UnixNano()),
				BuyOrderID:  buy.ID,
				SellOrderID: sell.ID,
				Price:       bestBid,
				Quantity:    tradeQty,
				Timestamp:   time.Now(),
			}

			trades = append(trades, t)

			sell.Quantity -= tradeQty
			buy.Quantity -= tradeQty

			if buy.Quantity == 0 {
				queue = queue[1:]
			}
		}

		if len(queue) == 0 {
			delete(ob.Bids, bestBid)
			ob.UpdateBidPrices()
		} else {
			ob.Bids[bestBid] = queue
		}
	}

	return trades
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
