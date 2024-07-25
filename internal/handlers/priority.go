package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/katerinapro/priority-api/internal/db"
	"github.com/katerinapro/priority-api/internal/models"
)

func GetPriority(c *gin.Context) {
	id := c.Param("id")
	priority, err := db.GetPriority(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, priority)
}

func CreatePriority(c *gin.Context) {
	var priority models.Priority
	if err := c.ShouldBindJSON(&priority); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.CreatePriority(&priority); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, priority)
}

func UpdatePriority(c *gin.Context) {
	id := c.Param("id")
	var priority models.Priority
	if err := c.ShouldBindJSON(&priority); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.UpdatePriority(id, &priority); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, priority)
}

func DeletePriority(c *gin.Context) {
	id := c.Param("id")
	if err := db.DeletePriority(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
