package main

import (
	"github.com/gin-gonic/gin"
	"github.com/katerinapro/priority-api/internal/db"
	"github.com/katerinapro/priority-api/internal/handlers"
)

func main() {
	db.Init()

	r := gin.Default()

	r.GET("v1/priorities", handlers.GetPriorities)
	r.GET("v1/priorities/:id", handlers.GetPriority)

	r.POST("v1/priorities", handlers.CreatePriority)
	r.PUT("v1/priorities/:id", handlers.UpdatePriority)
	r.DELETE("v1/priorities/:id", handlers.DeletePriority)

	r.Run(":8080")
}
