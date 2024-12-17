package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sayandas-sd/GinAuth/routes"
)

func main() {
	
		port := os.Getenv("PORT")

		if port == "" {
			port = "8080"
		}

		r := gin.Default()

		routes.authRouter(r)
		routes.userRouter(r)


		r.GET("/api-1", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "access granted for api-1",
			})
		})

		r.GET("/api-2", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "access granted for api-2",
			})
		})

		log.Printf("server is running on port %s...", port)

		if err := r.Run(":" + port); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	
}