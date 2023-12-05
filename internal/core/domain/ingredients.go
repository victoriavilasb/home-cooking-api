package domain

type Ingredient struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func MostPopularIngredientsInRecipes(recipes []Recipe) map[string]int {
	ingredientsRecipeCount := map[string]int{}
	for _, recipe := range recipes {
		for ingredient := range recipe.Ingredients {
			_, found := ingredientsRecipeCount[ingredient]
			if !found {
				ingredientsRecipeCount[ingredient] = 0
			}

			ingredientsRecipeCount[ingredient] += 1
		}
	}

	return ingredientsRecipeCount
}
