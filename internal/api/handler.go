package api

import (
	"net/http"
	"time"

	"gox_engine/internal/engine"
	"gox_engine/internal/order"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Engine *engine.Engine
}

func NewHandler(e *engine.Engine) *Handler {
	return &Handler{Engine: e}
}

// POST /orders
func (h *Handler) PlaceOrder(c *gin.Context) {
	var req struct {
		ID       string  `json:"id"`
		Type     string  `json:"type"`
		Side     string  `json:"side"`
		Price    float64 `json:"price"`
		Quantity int     `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	o := &order.Order{
		ID:        req.ID,
		Type:      order.OrderType(req.Type),
		Side:      order.Side(req.Side),
		Price:     req.Price,
		Quantity:  req.Quantity,
		Timestamp: time.Now(),
	}

	h.Engine.OrderChan <- o

	c.JSON(http.StatusOK, gin.H{"status": "order received"})
}

// DELETE /orders/:id
func (h *Handler) CancelOrder(c *gin.Context) {
	id := c.Param("id")

	h.Engine.CancelChan <- id

	c.JSON(http.StatusOK, gin.H{"status": "cancel request sent"})
}

// GET /events
func (h *Handler) GetEvents(c *gin.Context) {
	events := h.Engine.EventStore.GetAll()
	c.JSON(http.StatusOK, events)
}
