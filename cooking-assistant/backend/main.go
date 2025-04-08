package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"cooking-assistant/handlers"
	"cooking-assistant/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	storage.InitDB()

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/recipes", handlers.GetRecipes)
		api.GET("/recipes/:id", handlers.GetRecipe)
		api.POST("/recipes", handlers.CreateRecipe)
		api.PUT("/recipes/:id", handlers.UpdateRecipe)
		api.DELETE("/recipes/:id", handlers.DeleteRecipe)
	}

	frontendPath, err := filepath.Abs("../frontend")
	if err != nil {
		log.Fatal("Failed to get absolute path to frontend:", err)
	}

	r.Static("/static", filepath.Join(frontendPath, "static"))
	r.Static("/images", filepath.Join(frontendPath, "images"))

	r.GET("/", serveFrontend(filepath.Join(frontendPath, "index.html")))
	r.GET("/recipes", serveFrontend(filepath.Join(frontendPath, "recipes.html")))
	r.GET("/recipe-detail", serveFrontend(filepath.Join(frontendPath, "recipe-detail.html")))
	r.GET("/add-recipe", serveFrontend(filepath.Join(frontendPath, "add-recipe.html")))

	r.NoRoute(serveFrontend(filepath.Join(frontendPath, "index.html")))

	log.Println("Server is running on http://localhost:8080")
	log.Println("SQLite database file: backend/recipes.db")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func serveFrontend(file string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.File(file)
	}
}
