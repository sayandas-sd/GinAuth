package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sayandas-sd/GinAuth/controllers"
	"github.com/sayandas-sd/GinAuth/middleware"
)

func userRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate)
	incomingRoutes.GET("/users", controllers.GetUsers)
	incomingRoutes.GET("/users/:user_id", controllers.GetUser)
}
