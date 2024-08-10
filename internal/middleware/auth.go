package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckMiddleware(c *gin.Context) {
	headers := c.GetHeader("Authorization")
	fmt.Println(headers)

	c.Next()
}
