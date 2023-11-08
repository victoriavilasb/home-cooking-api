package repository

import "time"

type Recipe struct {
	ID         int
	ExternalID string
	Name       string
	Yield      int32
	CookTime   time.Duration
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Ingredient struct {
	ID         int
	ExternalID string
	Name       string
}

type Grocery struct {
	ID            int
	ExternalID    string
	Name          string
	Ingredient    string
	Type          string
	PurchaseDate  time.Time
	DueDate       time.Time
	IsPerishable  bool
	QuantityValue float64
	QuantityUnit  string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Available     bool
}

type RecipeIngredient struct {
	RecipeID      int
	IngredientID  int
	QuantityValue float64
	QuantityUnit  string
}
