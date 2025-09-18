package routes

import (
	"net/http"
	"time"

	"github.com/Anurag/url-shortner/api/database"
	"github.com/Anurag/url-shortner/api/models"
	"github.com/gin-gonic/gin"
)

func EditURL(c *gin.Context) {
	shortID := c.Param("shortID")

	var body models.Request

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	val, err := database.Client.Get(database.Ctx, shortID).Result()

	if err != nil || val == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	err = database.Client.Set(database.Ctx, shortID, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "URL updated successfully"})
}
