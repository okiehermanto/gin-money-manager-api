package helper

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("invalid authorization header")
	}

	token := strings.TrimPrefix(
		authHeader,
		"Bearer ",
	)

	return token, nil
}
