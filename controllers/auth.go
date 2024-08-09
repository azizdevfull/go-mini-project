package controllers

import "github.com/gin-gonic/gin"

type AuthController struct {
}

func InitController() *AuthController {
	return &AuthController{}
}
func (a *AuthController) InitRoutes(router *gin.Engine) {
	routes := router.Group("/auth")
	routes.POST("/login", a.Nope())
	routes.POST("/register", a.Nope())
}

func (*AuthController) Nope() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "nope",
		})
		return
	}
}
