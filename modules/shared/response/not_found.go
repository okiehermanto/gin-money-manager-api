package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "todo not found",
	})
}
