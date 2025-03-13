package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var recipes = []Recipe{}

type Recipe struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Ingredients string `json:"ingredients"`
	Instructions string `json:"instructions"`
}

func main() {
	router := gin.Default()

	// Routes
	router.GET("/recipes", GetRecipes)
	router.GET("/recipes/:id", GetRecipeByID)
	router.POST("/recipes", CreateRecipe)
	router.PUT("/recipes/:id", UpdateRecipe)
	router.DELETE("/recipes/:id", DeleteRecipe)

	// Start the server
	router.Run(":8080")
}

func GetRecipes(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

func GetRecipeByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, recipe := range recipes {
		if recipe.ID == id {
			c.JSON(http.StatusOK, recipe)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Recipe not found"})
}

func CreateRecipe(c *gin.Context) {
	var newRecipe Recipe
	if err := c.ShouldBindJSON(&newRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	// Assign a new ID
	newRecipe.ID = len(recipes) + 1
	recipes = append(recipes, newRecipe)
	c.JSON(http.StatusCreated, newRecipe)
}

func UpdateRecipe(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedRecipe Recipe
	if err := c.ShouldBindJSON(&updatedRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	for index, recipe := range recipes {
		if recipe.ID == id {
			recipes[index] = updatedRecipe
			recipes[index].ID = id
			c.JSON(http.StatusOK, recipes[index])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Recipe not found"})
}

func DeleteRecipe(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for index, recipe := range recipes {
		if recipe.ID == id {
			recipes = append(recipes[:index], recipes[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Recipe deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Recipe not found"})
}
