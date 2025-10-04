package routes

import (
	"net/http"

	"github.com/Anurag/url-shortner/api/database"
	"github.com/gin-gonic/gin"
)

func GetByShortID(c *gin.Context) {
	shortID := c.Param("shortID")

	val, err := database.Client.Get(database.Ctx, shortID).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": val})
}

func RedirectByShortID(c *gin.Context) {
	shortID := c.Param("shortID")

	val, err := database.Client.Get(database.Ctx, shortID).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, val)
}
