package main

import (
	"github.com/Safwanseban/Kubern8s_Docker_Project/configs"
	"github.com/Safwanseban/Kubern8s_Docker_Project/routes"
	"github.com/gin-gonic/gin"
)

var env = configs.Envload()
var R = gin.Default()

func init() {
	configs.ConnectDB()
}
func main() {
	routes.Routes(R)

	R.Run(env.Port)
}
