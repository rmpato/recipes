package mappers

import (
	"encoding/json"
	"net/http"
	"scrapper/domain"
)

func GetBonappetitRecipesFromJson(res *http.Response) (domain.BonappetitScrap, error) {
	var bonappetitRecipes domain.BonappetitScrap
	decErr := json.NewDecoder(res.Body).
		Decode(&bonappetitRecipes)
	return bonappetitRecipes, decErr
}