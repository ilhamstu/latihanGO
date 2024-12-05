package main

import (
	"os"
	"restapi/be/controllers"
	"restapi/be/middlewares"
	"restapi/be/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	models.ConnectDatabase()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	router.POST("/login", controllers.LoginHandler)
	router.POST("/admin", controllers.RegisterUser)
	router.POST("/daftar", controllers.RegisterMahasiswa)
	router.POST("/refresh", controllers.RefreshTokenHandler)

	postRoutes := router.Group("/posts")
	postRoutes.Use(middlewares.AuthMiddleware()) // Lindungi rute ini
	{
		postRoutes.GET("", controllers.FindPost)
		postRoutes.GET("/:id", controllers.FindPostByID)
		postRoutes.POST("", controllers.StorePost)
		postRoutes.PUT("/:id", controllers.UpdatePost)
		postRoutes.DELETE("/:id", controllers.DeletePost)
	}
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router.Run(":" + port)
}
