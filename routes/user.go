package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sayandas-sd/GinAuth/controllers"
	"github.com/sayandas-sd/GinAuth/middleware"
)

func userRoutes(r *gin.Engine) {
	r.Use(middleware.Authenticate)
	r.GET("/", controllers.GetUsers)
	r.GET("/:user_id", controllers.GetUser)
}
