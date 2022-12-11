package routes

import (
	"github.com/Safwanseban/Kubern8s_Docker_Project/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(c *gin.Engine) {
	c.POST("/signup", controllers.Signup)
	c.POST("/login", controllers.Login)
	c.GET("/", controllers.Home)
}
