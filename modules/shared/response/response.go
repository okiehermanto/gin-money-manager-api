package response

import (
	"fmt"
	"gin-money-manager-api/modules/shared/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Data  any `json:"data"`
	Meta  any `json:"meta,omitempty"`
	Links any `json:"links,omitempty"`
}

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
	var jsonResponse JsonResponse

	if err != nil {
		ServerError(c, err.Error())
		return
	}
	pagination, ok := data.(dto.Pagination)

	if ok {
		path := c.Request.URL.Path

		jsonResponse = JsonResponse{
			Data: pagination.Data,
			Meta: map[string]any{
				"page":  pagination.Page,
				"limit": pagination.Limit,
				"total": pagination.Total,
			},
			Links: buildLinks(path, pagination.Page, pagination.Limit, (pagination.Total+pagination.Limit-1)/pagination.Limit),
		}
	} else {
		jsonResponse = JsonResponse{
			Data: data,
		}
	}

	c.JSON(http.StatusOK, jsonResponse)
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

func buildLinks(base string, page, perPage, lastPage int) map[string]string {
	makeURL := func(p int) string {
		return fmt.Sprintf("%s?page=%d&limit=%d", base, p, perPage)
	}

	links := map[string]string{
		"self":  makeURL(page),
		"first": makeURL(1),
		"last":  makeURL(lastPage),
		"prev":  "",
		"next":  "",
	}

	if page > 1 {
		links["prev"] = makeURL(page - 1)
	}

	if page < lastPage {
		links["next"] = makeURL(page + 1)
	}

	return links
}
