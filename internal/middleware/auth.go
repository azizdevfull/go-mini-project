package middleware

import (
	"fmt"
	"go-tutorial/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckMiddleware(c *gin.Context) {
	headers := c.GetHeader("Authorization")

	if headers == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
	token := strings.Split(headers, " ")
	if len(token) < 2 {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	data, err := utils.TokenCheck(token[1])
	fmt.Println(data)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	c.Next()
}
