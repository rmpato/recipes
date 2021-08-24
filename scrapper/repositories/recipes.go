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
			for _, desc := range preparationGroup.Steps[0].Description {
				_, isString := desc.(string)
				if !isString {

					for _, value := range desc.([]interface{}) {
						fmt.Println("VALUE:")
						fmt.Println(value)
						if len(value.(string)) > 10 {
							steps = steps + value.(string)
						}
					}
				}
			}
		}


		var ingredients []db.RecipeIngredient

		for _, ig := range r.IngredientGroups{
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
			Name:        r.Title,
			Steps:       steps,
			Ingredients: ingredients,
		}
		recipes = append(recipes, recipe)
	}

	fmt.Println("TOTAL RECIPES: ")
	fmt.Println(len(recipes))
}