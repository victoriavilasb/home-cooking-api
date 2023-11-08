package ports

import (
	"context"

	"github.com/victoriavilasb/home-cooking-api/internal/core/domain"
)

// quando inserimos receitas os ingredientes são populados automaticamente
type RepositoryProvider interface {
}

type HomeCookingServiceProvider interface {
	RetrieveRecommendations(ctx context.Context, filter *RecommendationFilter) ([]domain.Recommendation, error)
	RegisterGrocery(ctx context.Context, grocery domain.Grocery) error
	UpdateGrocery(ctx context.Context, grocery domain.Grocery) error
	DeleteGrocery(ctx context.Context, id string) error
	RetrieveGroceries(ctx context.Context, filter *GroceryFilter) ([]domain.Grocery, error)
	RegisterRecipe(ctx context.Context, recipe domain.Recipe) error
}
