package main

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/victoriavilasb/home-cooking-api/internal/core/domain"
	"github.com/victoriavilasb/home-cooking-api/internal/repository"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/home-cooking?sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	logger := zap.NewNop().Sugar()

	// conn, _ := d/.Conn(ctx)
	repo := repository.NewRepositoryClient(db, logger)

	iID := uuid.New().String()
	repo.InsertIngredients(ctx, domain.Ingredient{
		ID:   iID,
		Name: "Bacon",
	})

	if err := repo.UpsertGrocery(ctx, domain.Grocery{
		ID:           uuid.NewString(),
		Name:         "Bacon Seara",
		PurchaseDate: "2023/10/01",
		DueDate:      "2023/11/02",
		Ingredient:   iID,
		IsPerishable: true,
		Quantity: domain.Quantity{
			Value: 500,
			Type:  "grams",
		},
	}); err != nil {
		panic(err)
	}

	if err := repo.UpsertGrocery(ctx, domain.Grocery{
		ID:           uuid.NewString(),
		Name:         "Ovos Felizes",
		PurchaseDate: "2023/10/01",
		DueDate:      "2023/11/02",
		Ingredient:   "c6b91b7a-8a1d-4166-aeb4-3a7a9b5d03bb",
		IsPerishable: true,
		Quantity: domain.Quantity{
			Value: 12,
			Type:  "unit",
		},
	}); err != nil {
		panic(err)
	}

	if err := repo.UpsertGrocery(ctx, domain.Grocery{
		ID:           uuid.NewString(),
		Name:         "Macarrão Basila",
		PurchaseDate: "2023/10/01",
		DueDate:      "2023/11/02",
		Ingredient:   "0d52d24e-052e-43ec-a038-a732ca8c40aa",
		IsPerishable: true,
		Quantity: domain.Quantity{
			Value: 1000,
			Type:  "grams",
		},
	}); err != nil {
		panic(err)
	}

	if err := repo.InsertRecipe(ctx, domain.Recipe{
		ID:   uuid.New().String(),
		Name: "Macarrão carbonara",
		Ingredients: map[string]domain.Quantity{
			iID: {
				Value: 250,
				Type:  "grams",
			},
			"0d52d24e-052e-43ec-a038-a732ca8c40aa": {
				Value: 250,
				Type:  "grams",
			},
			"c6b91b7a-8a1d-4166-aeb4-3a7a9b5d03bb": {
				Value: 3,
				Type:  "unit",
			},
		},
	}); err != nil {
		panic(err)
	}

}
