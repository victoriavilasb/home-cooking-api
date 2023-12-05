package domain

import "time"

type Grocery struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	PurchaseDate string   `json:"purchase_date"`
	Ingredient   string   `json:"ingredient"`
	DueDate      string   `json:"due_date"`
	IsPerishable bool     `json:"is_perishable"`
	Quantity     Quantity `json:"quantity"`
}

type Quantity struct {
	Value float64 `json:"value"`
	Type  string  `json:"type"`
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
