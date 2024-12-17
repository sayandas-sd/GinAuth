package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sayandas-sd/GinAuth/controllers"
)

func authRouter(r *gin.Engine) {
	authGroup := r.Group("/api")
	{
		authGroup.POST("/signup", controller.SignUp)
		authGroup.POST("/login", controller.Login)
	}
}
