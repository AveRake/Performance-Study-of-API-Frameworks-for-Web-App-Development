package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/recipe-app/config"
	"github.com/recipe-app/handlers"
	"github.com/recipe-app/database"
	"github.com/recipe-app/auth"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	// Initialize database
	database.InitDB()

	// Initialize Gin router
	r := gin.Default()

	// Routes
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)

	// Protected routes
	authRoutes := r.Group("/recipes")
	authRoutes.Use(auth.JWTMiddleware())
	{
		authRoutes.GET("/", handlers.GetAllRecipes)
		authRoutes.GET("/:id", handlers.GetRecipeByID)
		authRoutes.POST("/", handlers.CreateRecipe)
		authRoutes.PUT("/:id", handlers.UpdateRecipe)
		authRoutes.DELETE("/:id", handlers.DeleteRecipe)
	}

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
		os.Exit(1)
	}
}
