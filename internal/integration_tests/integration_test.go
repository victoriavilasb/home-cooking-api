package integrationtests_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/victoriavilasb/home-cooking-api/internal/core/domain"
	"github.com/victoriavilasb/home-cooking-api/internal/core/service"
	"github.com/victoriavilasb/home-cooking-api/internal/repository"
	"go.uber.org/zap"
)

func Test_RetrieveRecommendations(t *testing.T) {
	ctx := context.Background()

	logger := zap.NewNop().Sugar()

	connection := "postgresql://postgres:postgres@localhost:5432/home-cooking?sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		db.Close()
		panic(err)
	}
	defer db.Close()

	repo := repository.NewRepositoryClient(db, logger)

	svc, err := service.NewHomeCookingService(repo, logger)
	if err != nil {
		panic(err)
	}

	err = svc.RegisterGrocery(ctx, domain.Grocery{
		ID:           "grocery-1",
		Name:         "Rice",
		Type:         "Food",
		PurchaseDate: "2023-08-01",
		Ingredient:   "arroz",
		DueDate:      "2023-08-30",
		IsPerishable: false,
		Quantity:     domain.Quantity{Value: 1000, Type: "g"},
	})
	assert.Nil(t, err)

	err = svc.RegisterGrocery(ctx, domain.Grocery{
		ID:           "grocery-2",
		Name:         "Flour",
		Type:         "Food",
		PurchaseDate: "2023-08-01",
		Ingredient:   "farinha",
		DueDate:      "2023-08-15",
		IsPerishable: false,
		Quantity:     domain.Quantity{Value: 500, Type: "g"},
	})
	assert.Nil(t, err)

	recipe := domain.Recipe{
		ID:          "recipe-1",
		Name:        "Arroz branco",
		Ingredients: map[string]domain.Quantity{"arroz": {Value: 100, Type: "g"}},
		Yield:       5,
		CookTime:    time.Minute.String(),
	}
	err = svc.RegisterRecipe(ctx, recipe)
	assert.Nil(t, err)

	expected := domain.Recommendation{
		Recipe: recipe,
		Reason: domain.Reason(domain.AllIngredientsAvailable),
		Type:   domain.CookingRecommendation,
	}

	t.Run("should retrieve recommendations", func(t *testing.T) {
		got, err := svc.RetrieveRecommendations(ctx, nil)
		assert.Nil(t, err)
		assert.Equal(t, len(got), 1)

		assert.Exactly(t, got[0].Recipe.Name, expected.Recipe.Name)
	})
}
