package ports

type RecommendationFilter struct {
	Type *string
}

type GroceryFilter struct {
	ID         *string
	Name       *string
	Type       *string
	Ingredient *string
	Available  *bool
}

type RecipeFilter struct {
	ID    *string
	Name  *string
	Yield *int32
}
