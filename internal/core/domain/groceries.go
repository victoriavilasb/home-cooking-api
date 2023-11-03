package domain

import "time"

type Grocery struct {
	ID           string
	Name         string
	Type         string
	PurchaseDate string
	Ingredient   string
	DueDate      string
	IsPerishable bool
	Quantity     Quantity
}

type Quantity struct {
	Value float64
	Type  string
}

// verifica ingredients não disponíveis nos mantimentos
func GroceriesExpired(groceries []Grocery) []Grocery {
	var expired []Grocery
	for _, grocery := range groceries {
		today := time.Now().Format(time.DateOnly)
		if grocery.DueDate <= today {
			expired = append(expired, grocery)
		}
	}

	return expired
}
