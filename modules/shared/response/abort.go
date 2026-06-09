package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Unauthorized(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": err.Error(),
	})
}

func Forbidden(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"message": "Forbidden",
	})
}
