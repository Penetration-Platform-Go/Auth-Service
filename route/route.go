package route

import (
	"github.com/Penetration-Platform-Go/Auth-Service/controller"
	"github.com/gin-gonic/gin"
)

// GetServer Return Gin Server
func GetServer() *gin.Engine {
	app := gin.Default()
	app.POST("/auth", controller.LogInHandler)
	return app
}
