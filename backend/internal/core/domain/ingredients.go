package domain

type Ingredient struct {
	ID   string
	Name string
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

	// ingredients e a quantidade de vezes que eles aparecem
	return ingredientsRecipeCount
}
