package helper

import (
	"gin-money-manager-api/modules/shared/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindAndValidate(c *gin.Context, body any) bool {
	if err := c.ShouldBindJSON(body); err != nil {
		response.Error(c, http.StatusBadRequest, GetValidationError(err))

		return false
	}

	return true
}
