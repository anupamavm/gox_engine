package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *Handler) *gin.Engine {
	r := gin.Default()

	r.POST("/orders", h.PlaceOrder)
	r.DELETE("/orders/:id", h.CancelOrder)
	r.GET("/events", h.GetEvents)

	return r
}
