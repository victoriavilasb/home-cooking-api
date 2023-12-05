package ports

import (
	"context"

	"github.com/victoriavilasb/home-cooking-api/internal/core/domain"
)

// quando inserimos receitas os ingredientes são populados automaticamente
type RepositoryProvider interface {
	UpsertGrocery(ctx context.Context, grocery domain.Grocery) error
	InsertRecipe(ctx context.Context, recipe domain.Recipe) error
	ListGroceries(ctx context.Context, filter GroceryFilter) ([]domain.Grocery, error)
	ListRecipes(ctx context.Context, filter RecipeFilter) ([]domain.Recipe, error)
	InsertIngredients(ctx context.Context, ingredient domain.Ingredient) error
	UpsertRecipeIngredient(ctx context.Context, recipeID string, ingredientID string, quantityValue float64, quantityUnit string) error
}

type HomeCookingServiceProvider interface {
	RetrieveRecommendations(ctx context.Context, filter *RecommendationFilter) ([]domain.Recommendation, error)
	RegisterGrocery(ctx context.Context, grocery domain.Grocery) error
	UpdateGrocery(ctx context.Context, id string, grocery domain.Grocery) error
	DeleteGrocery(ctx context.Context, id string) error
	RetrieveGroceries(ctx context.Context, filter *GroceryFilter) ([]domain.Grocery, error)
	RegisterRecipe(ctx context.Context, recipe domain.Recipe) error
	IngredientsToBuy(ctx context.Context) ([]string, error)
}
