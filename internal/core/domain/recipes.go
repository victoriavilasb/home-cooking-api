package domain

import "time"

type Recipe struct {
	Name        string
	Ingredients map[Grocery]Quantity
	Yield       int32
	CookTime    time.Duration
}
