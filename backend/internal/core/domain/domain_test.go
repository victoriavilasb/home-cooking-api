package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/victoriavilasb/home-cooking-api/internal/core/domain"
)

func TestMostPopularIngredientsInRecipes(t *testing.T) {
	t.Run("Only one ingredient in different recipes", func(t *testing.T) {
		recipes := []domain.Recipe{
			{
				ID:   "id-1",
				Name: "Ovo mexido",
				Ingredients: map[string]domain.Quantity{
					"ovo": {
						Value: 2,
						Type:  "unit",
					},
				},
			},
			{
				ID:   "id-2",
				Name: "Omelete",
				Ingredients: map[string]domain.Quantity{
					"ovo": {
						Value: 3,
						Type:  "unit",
					},
				},
			},
			{
				ID:   "id-2",
				Name: "Ovo cozido",
				Ingredients: map[string]domain.Quantity{
					"ovo": {
						Value: 3,
						Type:  "unit",
					},
				},
			},
		}

		expected := map[string]int{
			"ovo": 3,
		}

		got := domain.MostPopularIngredientsInRecipes(recipes)

		assert.Exactly(t, expected, got)
	})

	t.Run("Only different ingredients in recipes", func(t *testing.T) {
		recipes := []domain.Recipe{
			{
				ID:   "id-1",
				Name: "Ovo mexido",
				Ingredients: map[string]domain.Quantity{
					"ovo": {
						Value: 2,
						Type:  "unit",
					},
				},
			},
			{
				ID:   "id-2",
				Name: "Brocolis Cozido",
				Ingredients: map[string]domain.Quantity{
					"brocolis": {
						Value: 200,
						Type:  "grams",
					},
					"alho": {
						Value: 2,
						Type:  "unit",
					},
				},
			},
			{
				ID:   "id-3",
				Name: "Macarrão",
				Ingredients: map[string]domain.Quantity{
					"macarrao": {
						Value: 200,
						Type:  "grams",
					},
					"molho-de-tomate": {
						Value: 50,
						Type:  "grams",
					},
				},
			},
		}

		expected := map[string]int{
			"ovo":             1,
			"brocolis":        1,
			"alho":            1,
			"macarrao":        1,
			"molho-de-tomate": 1,
		}

		got := domain.MostPopularIngredientsInRecipes(recipes)

		assert.Exactly(t, expected, got)
	})

	t.Run("Same ingredients in different recipes", func(t *testing.T) {
		recipes := []domain.Recipe{
			{
				ID:   "id-1",
				Name: "Macarrão ao alho e oleo",
				Ingredients: map[string]domain.Quantity{
					"macarrao": {
						Value: 2,
						Type:  "unit",
					},
					"alho": {
						Value: 2,
						Type:  "unit",
					},
					"azeite": {
						Value: 1,
						Type:  "tablespoon",
					},
				},
			},
			{
				ID:   "id-2",
				Name: "Brocolis Cozido",
				Ingredients: map[string]domain.Quantity{
					"brocolis": {
						Value: 200,
						Type:  "grams",
					},
					"alho": {
						Value: 2,
						Type:  "unit",
					},
				},
			},
			{
				ID:   "id-3",
				Name: "Farofa de ovo",
				Ingredients: map[string]domain.Quantity{
					"farinha": {
						Value: 200,
						Type:  "grams",
					},
					"ovo": {
						Value: 2,
						Type:  "unit",
					},
				},
			},
			{
				ID:   "id-4",
				Name: "Macarrão",
				Ingredients: map[string]domain.Quantity{
					"macarrao": {
						Value: 200,
						Type:  "grams",
					},
					"molho-de-tomate": {
						Value: 50,
						Type:  "grams",
					},
				},
			},
		}

		expected := map[string]int{
			"ovo":             1,
			"brocolis":        1,
			"alho":            2,
			"farinha":         1,
			"macarrao":        2,
			"molho-de-tomate": 1,
			"azeite":          1,
		}

		got := domain.MostPopularIngredientsInRecipes(recipes)

		assert.Exactly(t, expected, got)
	})
}

func TestExpiredGroceries(t *testing.T) {
	t.Run("no groceries expired", func(t *testing.T) {
		groceries := []domain.Grocery{
			{
				DueDate: time.Now().AddDate(0, 0, 1).String(),
			},
			{
				DueDate: time.Now().AddDate(0, 0, 2).String(),
			},
		}

		got := domain.GroceriesExpired(groceries)
		assert.Equal(t, len(got), 0)
	})

	t.Run("some expired groceries found", func(t *testing.T) {
		groceries := []domain.Grocery{
			{
				DueDate: time.Now().AddDate(0, 0, -3).Format(time.DateOnly),
			},
			{
				DueDate: time.Now().AddDate(0, 0, -2).Format(time.DateOnly),
			},
			{
				DueDate: time.Now().AddDate(0, 0, 5).Format(time.DateOnly),
			},
		}

		got := domain.GroceriesExpired(groceries)
		assert.Equal(t, len(got), 2)
	})

	t.Run("groceries expired today", func(t *testing.T) {
		groceries := []domain.Grocery{
			{
				DueDate: time.Now().Format(time.DateOnly),
			},
			{
				DueDate: time.Now().Format(time.DateOnly),
			},
		}

		got := domain.GroceriesExpired(groceries)
		assert.Equal(t, len(got), 2)
	})

	t.Run("all groceries expired", func(t *testing.T) {
		groceries := []domain.Grocery{
			{
				DueDate: time.Now().AddDate(0, 0, -1).String(),
			},
			{
				DueDate: time.Now().AddDate(0, 0, -2).String(),
			},
			{
				DueDate: time.Now().AddDate(0, 0, -2).String(),
			},
			{
				DueDate: time.Now().AddDate(0, 0, -2).String(),
			},
		}

		got := domain.GroceriesExpired(groceries)
		assert.Equal(t, len(got), 4)
	})
}

func TestRecipeIngredientsAvailable(t *testing.T) {
	groceries := []domain.Grocery{
		{
			DueDate: time.Now().AddDate(0, 0, 5).Format(time.DateOnly),
			Quantity: domain.Quantity{
				Value: 3,
				Type:  "unit",
			},
			Ingredient: "ovo",
		},
		{
			DueDate: time.Now().AddDate(0, 0, 5).Format(time.DateOnly),
			Quantity: domain.Quantity{
				Value: 500,
				Type:  "grams",
			},
			Ingredient: "bacon",
		},
		{
			DueDate: time.Now().AddDate(0, 0, 5).Format(time.DateOnly),
			Quantity: domain.Quantity{
				Value: 500,
				Type:  "grams",
			},
			Ingredient: "macarrao",
		},
	}

	t.Run("all recipe ingredients available", func(t *testing.T) {
		recipe := domain.Recipe{
			ID:   "recipe",
			Name: "Macarrão carbonara para 2",
			Ingredients: map[string]domain.Quantity{
				"ovo": {
					Value: 1,
					Type:  "unit",
				},
				"bacon": {
					Value: 100,
					Type:  "grams",
				},
				"macarrao": {
					Value: 200,
					Type:  "grams",
				},
			},
		}

		got := domain.IsRecipeIngredientsAvailable(recipe, groceries)
		assert.True(t, got)
	})

	t.Run("no recipe ingredients available", func(t *testing.T) {
		recipe := domain.Recipe{
			ID:   "recipe",
			Name: "Macarrão carbonara para 8",
			Ingredients: map[string]domain.Quantity{
				"ovo": {
					Value: 6,
					Type:  "unit",
				},
				"bacon": {
					Value: 600,
					Type:  "grams",
				},
				"macarrao": {
					Value: 1000,
					Type:  "grams",
				},
			},
		}

		got := domain.IsRecipeIngredientsAvailable(recipe, groceries)
		assert.False(t, got)
	})

	t.Run("one recipe ingredient quantity greater than grocery quantity", func(t *testing.T) {
		recipe := domain.Recipe{
			ID:   "recipe",
			Name: "Macarrão carbonara para 5",
			Ingredients: map[string]domain.Quantity{
				"ovo": {
					Value: 2,
					Type:  "unit",
				},
				"bacon": {
					Value: 200,
					Type:  "grams",
				},
				"macarrao": {
					Value: 600,
					Type:  "grams",
				},
			},
		}

		got := domain.IsRecipeIngredientsAvailable(recipe, groceries)
		assert.False(t, got)
	})

	t.Run("more than one grocery with the same ingredient", func(t *testing.T) {
		recipe := domain.Recipe{
			ID:   "recipe",
			Name: "Pavlova",
			Ingredients: map[string]domain.Quantity{
				"ovo": {
					Value: 15,
					Type:  "unit",
				},
			},
		}

		groceries = []domain.Grocery{
			{
				Name:       "ovo feliz",
				Ingredient: "ovo",
				Quantity: domain.Quantity{
					Value: 10,
					Type:  "unit",
				},
			},
			{
				Name:       "ovo de granja",
				Ingredient: "ovo",
				Quantity: domain.Quantity{
					Value: 5,
					Type:  "unit",
				},
			},
		}
		got := domain.IsRecipeIngredientsAvailable(recipe, groceries)
		assert.True(t, got)
	})

}

// func TestRecipesAvailable(t *testing.T) {
// 	recipes := []domain.Recipe{
// 		{
// 			ID:   "recipe-1",
// 			Name: "Pavlova",
// 			Ingredients: map[string]domain.Quantity{
// 				"ovo": {
// 					Value: 15,
// 					Type:  "unit",
// 				},
// 			},
// 		},
// 		{
// 			ID:   "recipe-2",
// 			Name: "Pavlova",
// 			Ingredients: map[string]domain.Quantity{
// 				"ovo": {
// 					Value: 2,
// 					Type:  "unit",
// 				},
// 				"mararrao": {
// 					Value: 200,
// 					Type:  "grams",
// 				},
// 			},
// 		},
// 		{
// 			ID:   "recipe-3",
// 			Name: "Bruschetta",
// 			Ingredients: map[string]domain.Quantity{
// 				"pao": {
// 					Value: 1,
// 					Type:  "unit",
// 				},
// 				"tomate": {
// 					Value: 3,
// 					Type:  "unit",
// 				},
// 				"cebola": {
// 					Value: 1,
// 					Type:  "unit",
// 				},
// 				"alho": {
// 					Value: 1,
// 					Type:  "unit",
// 				},
// 			},
// 		},
// 	}
// }
