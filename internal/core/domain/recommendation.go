package domain

type Reason string

type RecommendationType string

var (
	ShoppingRecommendation RecommendationType = "shopping"
	CookingRecommendation  RecommendationType = "cooking"
)

type Recommendation struct {
	Recipe  Recipe
	Grocery []Grocery
	Reason  Reason
	Type    RecommendationType
}

var (
	IngredientsDueDateExpiring    string = "there are some ingredients expiring"
	AllIngredientsAvailable       string = "all ingredients of the recipe are available"
	AlmostAllIngredientsAvailable string = "almost all ingredients of the recipe are available"
)
