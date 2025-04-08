package storage

import (
	"cooking-assistant/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("recipes.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Recipe{})

	if DB.First(&models.Recipe{}).Error != nil {
		initialRecipes := []models.Recipe{
			{
				Title:       "Паста Карбонара",
				Description: "Классическая итальянская паста с соусом из яиц, пармезана и гуанчиале.",
				PrepTime:    "15 мин",
				CookTime:    "15 мин",
				Image:       "https://www.chefmarket.ru/blog/wp-content/uploads/2019/05/carbonara-recipe-e1557844142474.jpg",
				Ingredients: []string{
					"Спагетти - 400 г",
					"Гуанчиале или панчетта - 150 г",
					"Яичные желтки - 4 шт",
					"Пармезан - 50 г",
					"Черный перец - по вкусу",
					"Соль - по вкусу",
				},
				Instructions: []string{
					"Поставить воду для пасты, посолить.",
					"Нарезать панчетту тонкими ломтиками.",
					"Обжарить панчетту до хрустящего состояния.",
					"Приготовить спагетти al dente.",
					"Смешать желтки с тертым пармезаном и перцем.",
					"Смешать пасту с панчеттой, затем с соусом.",
				},
			},
			{
				Title:       "Салат Цезарь",
				Description: "Легендарный салат с курицей, крутонами и соусом Цезарь.",
				PrepTime:    "20 мин",
				CookTime:    "10 мин",
				Image:       "https://images.gastronom.ru/-UHzDgNx-m0MMa6OR0ilz2qP7MB0mKQeGceObc9jpck/pr:recipe-cover-image/g:ce/rs:auto:0:0:0/L2Ntcy9hbGwtaW1hZ2VzLzVhNzFhZGY1LTM3MTYtNDlmMy04NDNlLTAwMTg4MGNiM2E0OS5qcGc.webp",
				Ingredients: []string{
					"Куриное филе - 300 г",
					"Листья салата романо - 1 кочан",
					"Белый хлеб - 200 г",
					"Пармезан - 50 г",
					"Чеснок - 2 зубчика",
					"Оливковое масло - 50 мл",
					"Яйца - 2 шт",
					"Лимонный сок - 2 ст.л.",
					"Анчоусы - 4 шт",
					"Горчица - 1 ч.л.",
					"Соль, перец - по вкусу",
				},
				Instructions: []string{
					"Обжарить куриное филе до готовности.",
					"Нарезать хлеб кубиками, обжарить с чесноком.",
					"Приготовить соус: смешать желтки, лимонный сок, анчоусы, горчицу, масло.",
					"Порвать салат руками, добавить крутоны и курицу.",
					"Полить соусом, посыпать пармезаном.",
				},
			},
			{
				Title:       "Тирамису",
				Description: "Итальянский десерт из печенья савоярди, сыра маскарпоне и кофе.",
				PrepTime:    "30 мин",
				CookTime:    "0 мин",
				Image:       "https://static01.nyt.com/images/2017/04/05/dining/05COOKING-TIRAMISU1/05COOKING-TIRAMISU1-square640.jpg",
				Ingredients: []string{
					"Савоярди (дамские пальчики) - 200 г",
					"Маскарпоне - 500 г",
					"Яйца - 4 шт",
					"Сахар - 100 г",
					"Кофе эспрессо - 300 мл",
					"Какао-порошок - для посыпки",
					"Марсала (по желанию) - 50 мл",
				},
				Instructions: []string{
					"Приготовить крепкий кофе, остудить.",
					"Отделить желтки от белков.",
					"Взбить желтки с сахаром до белой массы.",
					"Добавить маскарпоне, аккуратно перемешать.",
					"Взбить белки в крепкую пену, добавить к смеси.",
					"Обмакнуть савоярди в кофе и выложить в форму.",
					"Слой печенья, слой крема, повторить.",
					"Посыпать какао, охладить 4-6 часов.",
				},
			},
		}
		DB.Create(&initialRecipes)
	}
}

func GetRecipes() []models.Recipe {
	var recipes []models.Recipe
	DB.Find(&recipes)
	return recipes
}

func GetRecipe(id uint) (*models.Recipe, error) {
	var recipe models.Recipe
	result := DB.First(&recipe, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &recipe, nil
}

func AddRecipe(recipe models.Recipe) {
	DB.Create(&recipe)
}

func UpdateRecipe(id uint, updatedRecipe models.Recipe) bool {
	result := DB.Model(&models.Recipe{}).Where("id = ?", id).Updates(updatedRecipe)
	return result.RowsAffected > 0
}

func DeleteRecipe(id uint) bool {
	result := DB.Delete(&models.Recipe{}, id)
	return result.RowsAffected > 0
}
