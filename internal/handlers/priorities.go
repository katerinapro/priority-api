package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/katerinapro/priority-api/internal/db"
)

func GetPriorities(c *gin.Context) {
	priorities, err := db.GetPriorities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, priorities)
}
