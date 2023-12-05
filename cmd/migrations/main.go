package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Importe o driver PostgreSQL
)

func main() {
	// Conecte-se ao PostgreSQL
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/home-cooking?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	createRecipesTable(db)
	createIngredientsTable(db)
	createRecipeIngredientsTable(db)
	createGroceriesTable(db)
	addExternalIDConstraints(db)

	fmt.Println("Migrações concluídas com sucesso.")
}

func createRecipesTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS recipes (
			id SERIAL PRIMARY KEY,
			external_id TEXT,
			name TEXT,
			yield INT,
			cook_time INTERVAL,
			created_at TIMESTAMPTZ,
			updated_at TIMESTAMPTZ
		)
	`)
	if err != nil {
		panic(err)
	}
}

func createIngredientsTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS ingredients (
			id SERIAL PRIMARY KEY,
			external_id TEXT,
			name TEXT
		)
	`)
	if err != nil {
		panic(err)
	}
}

func createRecipeIngredientsTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS recipe_ingredients (
			recipe_id TEXT,
			ingredient_id TEXT,
			PRIMARY KEY (recipe_id, ingredient_id),
			quantity_value FLOAT,
			quantity_unit TEXT
		)
	`)
	if err != nil {
		panic(err)
	}
}

func createGroceriesTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS groceries (
			id SERIAL PRIMARY KEY,
			external_id TEXT,
			name TEXT,
			type TEXT,
			ingredient TEXT,
			purchase_date TIMESTAMPTZ,
			due_date TIMESTAMPTZ,
			is_perishable BOOLEAN,
			quantity_value FLOAT,
			quantity_unit TEXT,
			available BOOLEAN,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
		)
	`)
	if err != nil {
		panic(err)
	}
}

func addExternalIDConstraints(db *sql.DB) {
	_, err := db.Exec(`
		ALTER TABLE ingredients
		ADD CONSTRAINT unique_external_id UNIQUE (external_id);
	`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		ALTER TABLE groceries
		ADD CONSTRAINT groceries_unique_external_id UNIQUE (external_id);
	`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		ALTER TABLE recipes
		ADD CONSTRAINT recipes_unique_external_id UNIQUE (external_id);
	`)
	if err != nil {
		panic(err)
	}
}
