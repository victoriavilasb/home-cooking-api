package repository

import (
	"time"
)

type Recipe struct {
	ID         int           `json:"id" db:"id"`
	ExternalID string        `json:"external_id" db:"external_id"`
	Name       string        `json:"name" db:"name"`
	Yield      int32         `json:"yield" db:"yield"`
	CookTime   time.Duration `json:"cook_time" db:"cook_time"`
	CreatedAt  time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at" db:"updated_at"`
}

type Ingredient struct {
	ID         int    `json:"id" db:"id"`
	ExternalID string `json:"external_id" db:"external_id"`
	Name       string `json:"name" db:"name"`
}

type Grocery struct {
	ID            int       `json:"id" db:"id"`
	ExternalID    string    `json:"external_id" db:"external_id"`
	Name          string    `json:"name" db:"name"`
	Ingredient    string    `json:"ingredient" db:"ingredient"`
	Type          string    `json:"type" db:"type"`
	PurchaseDate  time.Time `json:"purchase_date" db:"purchase_date"`
	DueDate       time.Time `json:"due_date" db:"due_date"`
	IsPerishable  bool      `json:"is_perishable" db:"is_perishable"`
	QuantityValue float64   `json:"quantity_value" db:"quantity_value"`
	QuantityUnit  string    `json:"quantity_unit" db:"quantity_unit"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	Available     bool      `json:"available" db:"available"`
}

type RecipeIngredient struct {
	RecipeID      int     `json:"recipe_id" db:"recipe_id"`
	IngredientID  int     `json:"ingredient_id" db:"ingredient_id"`
	QuantityValue float64 `json:"quantity_value" db:"quantity_value"`
	QuantityUnit  string  `json:"quantity_unit" db:"quantity_unit"`
}
