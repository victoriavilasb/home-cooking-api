package domain

import "time"

type Recipe struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Ingredients map[string]Quantity `json:"ingredients"`
	Yield       int32               `json:"yield"`
	CookTime    time.Duration       `json:"cook_time"`
}

func IsRecipeIngredientsAvailable(recipe Recipe, groceries []Grocery) bool {
	ingredientGroceryQuantityMapping := map[string]Quantity{}
	for _, grocery := range groceries {
		_, found := ingredientGroceryQuantityMapping[grocery.Ingredient]
		if !found {
			ingredientGroceryQuantityMapping[grocery.Ingredient] = Quantity{
				Value: 0,
				Type:  grocery.Quantity.Type,
			}
		}

		accValue := ingredientGroceryQuantityMapping[grocery.Ingredient].Value
		accValue += grocery.Quantity.Value

		ingredientGroceryQuantityMapping[grocery.Ingredient] = Quantity{
			Value: accValue,
			Type:  grocery.Quantity.Type,
		}
	}

	for ingredient, recipeQuantity := range recipe.Ingredients {
		quantity, found := ingredientGroceryQuantityMapping[ingredient]
		if !found {
			return false
		}

		// ainda não temos conversão de unidades
		if recipeQuantity.Value > quantity.Value {
			return false
		}
	}

	return true
}

// // retorna um array com todos os ingredientes faltantes
func MissingRecipeIngredients(recipe Recipe, groceries []Grocery) []string {
	missing := []string{}

	for ingredient, quantity := range recipe.Ingredients {
		found := false
		for _, grocery := range groceries {
			if grocery.Ingredient == ingredient {
				if grocery.Quantity.Value >= quantity.Value {
					found = true
					break
				}
			}
		}
		if !found {
			missing = append(missing, ingredient)
		}
	}

	return missing
}

// retorna as receitas cujos ingredientes estão disponíveis
func RecipesAvailable(recipes []Recipe, groceries []Grocery) []string {
	var recipesAvailable []string
	for _, recipe := range recipes {
		available := IsRecipeIngredientsAvailable(recipe, groceries)
		if !available {
			continue
		}

		recipesAvailable = append(recipesAvailable, recipe.Name)
	}

	return recipesAvailable
}
