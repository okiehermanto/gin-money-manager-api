package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "todo not found",
	})
}

func ModelNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": fmt.Sprintf("%s not found", message),
	})
}
