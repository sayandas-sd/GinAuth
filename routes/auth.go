package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sayandas-sd/GinAuth/controllers"
)

func authRouter(r *gin.Engine) {
	authGroup := r.Group("/api")
	{
		authGroup.POST("/signup", controllers.SignUp)
		authGroup.POST("/login", controllers.Login)
	}
}
