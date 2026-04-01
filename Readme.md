# 🚀 Order Matching Engine (Go)

A high-performance in-memory **Order Matching Engine** inspired by modern electronic exchanges like NASDAQ.

This project demonstrates core concepts used in **quant trading systems**, including:

- Price-time priority matching
- Limit and market orders
- Partial fills
- In-memory low-latency design

---

## 🧠 Features

- ✅ Limit & Market Orders
- ✅ Price-Time Priority Matching
- ✅ Partial Order Execution
- ✅ In-Memory Order Book
- ✅ FIFO Queue per Price Level
- ✅ Trade Generation

---

## ⚙️ Architecture

```
[Client Orders]
      ↓
[Order Book]
      ↓
[Matching Engine]
      ↓
[Trade Execution]
```

### Components:

- `order` → Order models
- `orderbook` → Bid/Ask storage
- `matching` → Core matching logic
- `trade` → Trade execution model

---

## 📚 Matching Logic

### Price Priority

- Best price is matched first
  - Highest bid
  - Lowest ask

### Time Priority

- Orders at the same price are matched FIFO

---

## 🔄 Example

```
Sell Orders | Price Qty
100           10
101           5
```

Incoming Buy (Market, Qty=12):

```
Execution:
- 10 @ 100
- 2 @ 101
```

---
