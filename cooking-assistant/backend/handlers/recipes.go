package handlers

import (
	"net/http"
	"strconv"

	"cooking-assistant/models"
	"cooking-assistant/storage"

	"github.com/gin-gonic/gin"
)

func GetRecipes(c *gin.Context) {
	recipes := storage.GetRecipes()
	c.JSON(http.StatusOK, recipes)
}

func GetRecipe(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID рецепта"})
		return
	}

	recipe, err := storage.GetRecipe(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Рецепт не найден"})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

func CreateRecipe(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storage.AddRecipe(recipe)
	c.JSON(http.StatusCreated, recipe)
}

func UpdateRecipe(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID рецепта"})
		return
	}

	var updatedRecipe models.Recipe
	if err := c.ShouldBindJSON(&updatedRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRecipe.ID = uint(id)

	if success := storage.UpdateRecipe(uint(id), updatedRecipe); !success {
		c.JSON(http.StatusNotFound, gin.H{"error": "Рецепт не найден"})
		return
	}

	c.JSON(http.StatusOK, updatedRecipe)
}

func DeleteRecipe(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID рецепта"})
		return
	}

	if success := storage.DeleteRecipe(uint(id)); !success {
		c.JSON(http.StatusNotFound, gin.H{"error": "Рецепт не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Рецепт успешно удален"})
}
