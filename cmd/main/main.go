package main

import (
	"fmt"

	"gox_engine/internal/api"
	"gox_engine/internal/engine"
)

func main() {
	engine := engine.NewEngine()
	engine.Start()

	handler := api.NewHandler(engine)
	router := api.SetupRouter(handler)

	fmt.Println("🚀 Server running on http://localhost:8080")

	router.Run(":8080")
}
