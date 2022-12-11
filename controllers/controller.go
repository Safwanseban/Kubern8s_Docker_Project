package controllers

import (
	"github.com/Safwanseban/Kubern8s_Docker_Project/configs"
	"github.com/Safwanseban/Kubern8s_Docker_Project/models"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(404, gin.H{
			"err": err.Error(),
		})
	}
	record := configs.Db.Create(&user)
	if record.Error != nil {
		c.JSON(500, gin.H{
			"err": record.Error.Error(),
		})
	}
	c.JSON(200, gin.H{
		"status": true,
	})

}
func Login(c *gin.Context) {
	var user models.User
	var login struct {
		Email    string
		Password string
	}
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(404, gin.H{
			"err": err.Error(),
		})
	}
	configs.Db.Raw("select email,password from users where email=? ", login.Email).Scan(&user)
	if user.Email == login.Email && user.Password == login.Password {
		c.JSON(200, gin.H{
			"status": true,
		})
	} else {
		c.JSON(401, gin.H{
			"status ": false,
		})
	}

}
func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "hai",
	})
}
