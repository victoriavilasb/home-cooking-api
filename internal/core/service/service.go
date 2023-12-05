package service

import (
	"context"
	"errors"

	"github.com/victoriavilasb/home-cooking-api/internal/core/domain"
	"github.com/victoriavilasb/home-cooking-api/internal/core/ports"
	"go.uber.org/zap"
)

type HomeCookingService struct {
	repository ports.RepositoryProvider
	logger     *zap.SugaredLogger
}

func NewHomeCookingService(repository ports.RepositoryProvider, logger *zap.SugaredLogger) (*HomeCookingService, error) {
	if logger == nil {
		err := errors.New("logger is required")
		return nil, err
	}

	if repository == nil {
		err := errors.New("repository cannot be nil")
		logger.Errorf("Missing required dependency: %s", err)
		return nil, err
	}

	svc := &HomeCookingService{
		repository: repository,
		logger:     logger,
	}

	return svc, nil
}

func (s *HomeCookingService) RetrieveRecommendations(ctx context.Context, filter *ports.RecommendationFilter) ([]domain.Recommendation, error) {
	groceries, err := s.repository.ListGroceries(ctx, ports.GroceryFilter{})
	if err != nil {
		return nil, err
	}

	recipes, err := s.repository.ListRecipes(ctx, ports.RecipeFilter{})
	if err != nil {
		return nil, err
	}

	recommendations := domain.AvailableRecipes(recipes, groceries)

	recipesMap := map[string]domain.Recipe{}
	for _, recipe := range recipes {
		recipesMap[recipe.Name] = recipe
	}

	availableRecommendations := make([]domain.Recommendation, len(recommendations))
	for i, rec := range recommendations {
		availableRecommendations[i] = domain.Recommendation{
			Recipe: recipesMap[rec],
			Reason: domain.Reason(domain.AllIngredientsAvailable),
			Type:   domain.CookingRecommendation,
		}
	}

	return availableRecommendations, nil
}

func (s *HomeCookingService) IngredientsToBuy(ctx context.Context) ([]string, error) {
	groceries, err := s.repository.ListGroceries(ctx, ports.GroceryFilter{})
	if err != nil {
		return nil, err
	}

	recipe, err := s.repository.ListRecipes(ctx, ports.RecipeFilter{})
	if err != nil {
		return nil, err
	}

	ingredients := domain.IngredientsNotAvailable(recipe, groceries)

	return ingredients, nil
}

func (s *HomeCookingService) RegisterGrocery(ctx context.Context, grocery domain.Grocery) error {
	if err := s.repository.InsertIngredients(ctx, domain.Ingredient{
		ID:   grocery.Ingredient,
		Name: grocery.Ingredient,
	}); err != nil {
		return err
	}

	return s.repository.UpsertGrocery(ctx, grocery)
}

func (s *HomeCookingService) UpdateGrocery(ctx context.Context, id string, grocery domain.Grocery) error {

	return nil
}

func (s *HomeCookingService) DeleteGrocery(ctx context.Context, id string) error {
	return nil
}

func (s *HomeCookingService) RetrieveGroceries(ctx context.Context, filter *ports.GroceryFilter) ([]domain.Grocery, error) {
	groceries := []domain.Grocery{}

	return groceries, nil
}

func (s *HomeCookingService) RegisterRecipe(ctx context.Context, recipe domain.Recipe) error {
	return s.repository.InsertRecipe(ctx, recipe)
}
