package domain

import "time"

type Recipe struct {
	ID          string
	Name        string
	Ingredients map[string]Quantity
	Yield       int32
	CookTime    time.Duration
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
// func MissingRecipeIngredients(recipe Recipe, groceries []Grocery) map[string]Quantity {
// 	ingredientGroceryMapping := map[string]Grocery{}
// 	missingIngredients := map[string]Quantity{}
// 	for _, grocery := range groceries {
// 		ingredientGroceryMapping[grocery.Ingredient] = grocery
// 	}

// 	for ingredient, recipeQuantity := range recipe.Ingredients {
// 		grocery, found := ingredientGroceryMapping[ingredient]
// 		if !found {
// 			missingIngredients[ingredient] = recipeQuantity
// 			continue
// 		}

// 		// ainda não temos conversão de unidades
// 		if recipeQuantity.Value > grocery.Quantity.Value {
// 			missingIngredients[ingredient] = Quantity{
// 				Value: ,
// 			}
// 			return false
// 		}

// 		today := time.Now().Format(time.DateOnly)
// 		if grocery.DueDate < today {
// 			return false
// 		}
// 	}

// 	return true

// }

// retorna as receitas cujos ingredientes estão disponíveis
// func RecipesAvailable([]Recipe, []Grocery) []string {

// }
