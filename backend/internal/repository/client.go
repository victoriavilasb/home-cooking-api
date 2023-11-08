package repository

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq" // Importe o driver PostgreSQL
	"github.com/victoriavilasb/home-cooking-api/internal/core/domain"
	"github.com/victoriavilasb/home-cooking-api/internal/core/ports"
	"go.uber.org/zap"
)

type RepositoryClient struct {
	DB     *sql.DB
	Logger *zap.SugaredLogger
}

func NewRepositoryClient(db *sql.DB, logger *zap.SugaredLogger) RepositoryClient {
	return RepositoryClient{
		DB:     db,
		Logger: logger,
	}
}

func (r *RepositoryClient) UpsertGrocery(ctx context.Context, grocery domain.Grocery) error {
	query := `
		INSERT INTO groceries (external_id, name, type, ingredient, purchase_date, due_date, is_perishable, quantity_value, quantity_unit, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW())
		ON CONFLICT (external_id) DO UPDATE
		SET name = $2, type = $3, ingredient = $4, purchase_date = $5, due_date = $6, is_perishable = $7, quantity_value = $8, quantity_unit = $9, updated_at = NOW()
	`

	// Executar a consulta SQL
	_, err := r.DB.Exec(query, grocery.ID, grocery.Name, grocery.Type, grocery.Ingredient, grocery.PurchaseDate, grocery.DueDate, grocery.IsPerishable, grocery.Quantity.Value, grocery.Quantity.Type)
	return err
}

func (r *RepositoryClient) DeleteGrocery(ctx context.Context, externalID string) error {
	query := `
		UPDATE groceries
		SET available = false
		WHERE external_id = $1
	`

	_, err := r.DB.Exec(query, externalID)
	return err
}

func (r *RepositoryClient) InsertRecipe(ctx context.Context, recipe domain.Recipe) error {
	conn, _ := r.DB.Conn(ctx)

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Inserir a receita
	query := `
		INSERT INTO recipes (external_id, name, yield, cook_time, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id
	`
	var recipeID string
	err = tx.QueryRowContext(ctx, query, recipe.ID, recipe.Name, recipe.Yield, recipe.CookTime.String()).Scan(&recipeID)
	if err != nil {
		return err
	}

	for ingredientID, quantity := range recipe.Ingredients {
		query = `
			INSERT INTO recipe_ingredients (recipe_id, ingredient_id, quantity_value, quantity_unit)
			VALUES ($1, $2, $3, $4)
		`
		_, err = tx.ExecContext(ctx, query, recipeID, ingredientID, quantity.Value, quantity.Type)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *RepositoryClient) ListGroceries(ctx context.Context, filter ports.GroceryFilter) ([]domain.Grocery, error) {
	// Construir a consulta SQL base
	query := `
		SELECT id, external_id, name, type, ingredient, purchase_date, due_date, is_perishable, quantity_value, quantity_unit, available, created_at, updated_at
		FROM groceries
		WHERE 1=1
	`

	// Construir as cláusulas SQL com base nos critérios do filtro
	args := []interface{}{}
	if filter.ID != nil {
		query += " AND id = $1"
		args = append(args, *filter.ID)
	}
	if filter.Name != nil {
		query += " AND name = $2"
		args = append(args, *filter.Name)
	}
	if filter.Type != nil {
		query += " AND type = $3"
		args = append(args, *filter.Type)
	}
	if filter.Ingredient != nil {
		query += " AND ingredient = $4"
		args = append(args, *filter.Ingredient)
	}
	if filter.Available != nil {
		query += " AND available = $5"
		args = append(args, *filter.Available)
	}

	// Executar a consulta SQL
	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groceries []domain.Grocery
	for rows.Next() {
		var grocery Grocery
		err := rows.Scan(
			&grocery.ID,
			&grocery.ExternalID,
			&grocery.Name,
			&grocery.Type,
			&grocery.Ingredient,
			&grocery.PurchaseDate,
			&grocery.DueDate,
			&grocery.IsPerishable,
			&grocery.QuantityValue,
			&grocery.QuantityUnit,
			&grocery.Available,
			&grocery.CreatedAt,
			&grocery.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		groceries = append(groceries, domain.Grocery{
			ID:           grocery.ExternalID,
			Name:         grocery.Name,
			Type:         grocery.Type,
			PurchaseDate: grocery.PurchaseDate.Format(time.DateOnly), // Formatar para string
			DueDate:      grocery.DueDate.Format(time.DateOnly),      // Formatar para string
			IsPerishable: grocery.IsPerishable,
			Quantity: domain.Quantity{
				Value: grocery.QuantityValue,
				Type:  grocery.QuantityUnit,
			},
		})
	}

	return groceries, nil
}

func (r *RepositoryClient) ListRecipes(ctx context.Context, filter ports.RecipeFilter) ([]domain.Recipe, error) {
	// Construir a consulta SQL base
	query := `
		SELECT id, external_id, name, yield, cook_time, created_at, updated_at
		FROM recipes
		WHERE 1=1
	`

	// Construir as cláusulas SQL com base nos critérios do filtro
	args := []interface{}{}
	if filter.ID != nil {
		query += " AND external_id = $1"
		args = append(args, *filter.ID)
	}
	if filter.Name != nil {
		query += " AND name = $2"
		args = append(args, *filter.Name)
	}
	if filter.Yield != nil {
		query += " AND yield = $3"
		args = append(args, *filter.Yield)
	}

	// Executar a consulta SQL
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []domain.Recipe
	for rows.Next() {
		var recipe Recipe
		err := rows.Scan(
			&recipe.ID,
			&recipe.ExternalID,
			&recipe.Name,
			&recipe.Yield,
			&recipe.CookTime,
			&recipe.CreatedAt,
			&recipe.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, domain.Recipe{
			ID:       recipe.ExternalID,
			Name:     recipe.Name,
			Yield:    recipe.Yield,
			CookTime: recipe.CookTime,
		})
	}

	return recipes, nil
}

func (r *RepositoryClient) InsertIngredients(ctx context.Context, ingredient domain.Ingredient) error {
	query := `
		INSERT INTO ingredients (external_id, name)
		VALUES ($1, $2)
	`

	_, err := r.DB.Exec(query, ingredient.ID, ingredient.Name)
	if err != nil {
		return err
	}

	return nil
}
