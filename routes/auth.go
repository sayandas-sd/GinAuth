package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sayandas-sd/GinAuth/controllers"
)

func authRouter(r *gin.Engine) {
	incomingRoutes := r.Group("/api")
	incomingRoutes.POST("/signup", controller.SignUp)
	incomingRoutes.POST("/login", controller.Login)

}
