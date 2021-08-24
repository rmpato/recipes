package domain

type BonappetitScrap struct {
	Recipes []BonappetitRecipe `json:"items"`
}
type BonappetitRecipe struct {
	IngredientGroups  []IngredientGroup `json:"ingredientGroups"`
	PreparationGroups []struct {
		Hed   string `json:"hed"`
		Steps []struct {
			Description []interface{} `json:"description"`
		} `json:"steps"`
		MicroSteps []interface{} `json:"microSteps"`
	} `json:"preparationGroups"`
	Title string `json:"hed"`
	Id string `json:"id"`
}

type IngredientGroup struct {
	//Title string `json:"hed"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name string `json:"name"`
	Description string `json:"description"`
}
