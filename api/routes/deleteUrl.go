package routes

import (
	"net/http"

	"github.com/Anurag/url-shortner/api/database"
	"github.com/gin-gonic/gin"
)

func DeleteURL(c *gin.Context) {
	shortID := c.Param("shortID")

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Del(database.Ctx, shortID).Result()
	if err != nil || val == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "URL deleted successfully",
	})
}
