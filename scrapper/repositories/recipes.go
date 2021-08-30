package repositories

import (
	"fmt"
	"scrapper/db"
	"scrapper/domain"
)

func Save(bonappetitRecipes domain.BonappetitScrap) {
	var recipes []db.Recipe

	for _, r := range bonappetitRecipes.Recipes {
		var steps string
		for _, preparationGroup := range r.PreparationGroups {

			for _, prepStep := range preparationGroup.Steps {
				for _, desc := range prepStep.Description {
					_, isString := desc.(string)
					if !isString {

						for _, value := range desc.([]interface{}) {

							_, valueIsString := value.(string)

							if valueIsString {
								if len(value.(string)) > 3 {
									steps = steps + value.(string)
								}
							}
						}
					}
				}
			}
		}

		var ingredients []db.RecipeIngredient

		for _, ig := range r.IngredientGroups {
			for _, baIngredient := range ig.Ingredients {
				ingredient := db.RecipeIngredient{
					Name:        baIngredient.Name,
					Description: baIngredient.Name,
				}

				ingredients = append(ingredients, ingredient)
			}
		}

		recipe := db.Recipe{
			ExternalId:  r.Id,
			Name:        string(r.Title),
			Steps:       steps,
			Ingredients: ingredients,
		}
		recipes = append(recipes, recipe)
	}

	recipesDb := db.Init()
	for _, recipe := range recipes {
		var recipeExists db.Recipe
		result := recipesDb.Find(&recipeExists, "external_id = ?", recipe.ExternalId)

		if result.RowsAffected == 0 {
			recipesDb.Create(&recipe)
			fmt.Println("Insertando '"+ recipe.Name)
		}
	}
}
