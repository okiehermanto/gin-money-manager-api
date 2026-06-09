package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"message": message,
	})
}

func ServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": message,
	})
}

func Success(c *gin.Context, data any, err error) {
	if err != nil {
		ServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func Created(c *gin.Context, data any, err error) {
	if err != nil {
		ServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": data,
	})
}

func Updated(c *gin.Context, data any, err error) {
	if err != nil {
		ServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusUpgradeRequired, gin.H{
		"data": data,
	})
}
