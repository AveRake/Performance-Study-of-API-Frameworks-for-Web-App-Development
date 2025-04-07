package models

type Recipe struct {
	ID           int      `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	PrepTime     string   `json:"prepTime"`
	CookTime     string   `json:"cookTime"`
	Image        string   `json:"image"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
}
