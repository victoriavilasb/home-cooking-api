package domain

type Reason string

type RecommendationType string

var RuleRecommendationExpiration = 2

var (
	ShoppingRecommendation RecommendationType = "shopping"
	CookingRecommendation  RecommendationType = "cooking"
)

type Recommendation struct {
	Recipe Recipe             `json:"recipe"`
	Reason Reason             `json:"reason"`
	Type   RecommendationType `json:"type"`
}

var (
	IngredientsDueDateExpiring    string = "there are some ingredients expiring"
	AllIngredientsAvailable       string = "all ingredients of the recipe are available"
	AlmostAllIngredientsAvailable string = "almost all ingredients of the recipe are available"
)

// verifica ingredients não disponíveis nos mantimentos
func IngredientsNotAvailable(recipes []Recipe, groceries []Grocery) []string {
	return nil
}

func AvailableRecipes(recipes []Recipe, groceries []Grocery) []string {
	available := []string{}

	for _, recipe := range recipes {
		requiredIngredients := map[string]float64{}
		for ingredient, quantity := range recipe.Ingredients {
			requiredIngredients[ingredient] = quantity.Value
		}

		for _, grocery := range groceries {
			if _, ok := requiredIngredients[grocery.Ingredient]; ok {
				if grocery.Quantity.Value < requiredIngredients[grocery.Ingredient] {
					continue
				}
				grocery.Quantity.Value -= requiredIngredients[grocery.Ingredient]
				if grocery.Quantity.Value > 0 {
					available = append(available, recipe.Name)
				}
			}
		}
	}

	return available
}
