package repository

import (
	"database/sql"
	"news/models"
)

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: db,
	}
}

// Create creates a new category
func (r *CategoryRepository) Create(category *models.Category) error {
	query := `
		INSERT INTO categories (name, slug)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`

	err := r.DB.QueryRow(
		query,
		category.Name,
		category.Slug,
	).Scan(
		&category.ID,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

// GetByID returns a category by ID
func (r *CategoryRepository) GetByID(id int64) (*models.Category, error) {
	var category models.Category

	query := `
		SELECT id, name, slug, created_at, updated_at
		FROM categories
		WHERE id = $1
	`

	err := r.DB.QueryRow(query, id).Scan(
		&category.ID,
		&category.Name,
		&category.Slug,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

// GetAll returns all categories
func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	query := `
		SELECT id, name, slug, created_at, updated_at
		FROM categories
		ORDER BY id ASC
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		var category models.Category

		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Slug,
			&category.CreatedAt,
			&category.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
