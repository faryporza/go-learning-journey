package routes

import (
	"github.com/farypor/go-journey/go-gin-fiber/08-gin-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/users", controllers.GetUsers)
		api.POST("/users", controllers.CreateUser)
	}
}
