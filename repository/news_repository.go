package repository

import (
	"database/sql"
	"news/models"
)

type NewsRepository struct {
	DB *sql.DB
}

func NewNewsRepository(db *sql.DB) *NewsRepository {
	return &NewsRepository{
		DB: db,
	}
}

// Create creates a new news article
func (r *NewsRepository) Create(news *models.News) error {
	query := `
    INSERT INTO news (
        title,
        slug,
        summary,
        content,
        image_url,
        status,
        category_id,
        author_id
    )
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id, created_at, updated_at
`
	err := r.DB.QueryRow(query, news.Title, news.Slug, news.Summary, news.Content, news.ImageURL, news.Status, news.CategoryID, news.AuthorID).Scan(&news.ID, &news.CreatedAt, &news.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// GetByID returns a news article by ID
func (r *NewsRepository) GetByID(id int64) (*models.News, error) {
	var news models.News

	query := `SELECT * FROM news WHERE id = $1`

	err := r.DB.QueryRow(query, id).Scan(&news.ID, &news.Title, &news.Slug, &news.Summary, &news.Content, &news.ImageURL, &news.Status, &news.CategoryID, &news.AuthorID, &news.CreatedAt, &news.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &news, nil
}

// GetAll returns all news articles
func (r *NewsRepository) GetAll() ([]models.News, error) {
	query := `SELECT id, title, slug, summary, content, image_url, status, category_id, author_id, created_at, updated_at
        FROM news
        ORDER BY id DESC
    `
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var newsList []models.News

	for rows.Next() {
		var news models.News

		err := rows.Scan(&news.ID, &news.Title, &news.Slug, &news.Summary, &news.Content, &news.ImageURL, &news.Status, &news.CategoryID, &news.AuthorID, &news.CreatedAt, &news.UpdatedAt)
		if err != nil {
			return nil, err
		}

		newsList = append(newsList, news)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return newsList, nil
}

// Update updates an existing news article.
func (r *NewsRepository) Update(news *models.News) error {
	query := `UPDATE news SET title = $1, slug = $2, summary = $3, content = $4, image_url = $5, status = $6, category_id = $7, author_id = $8, updated_at = CURRENT_TIMESTAMP WHERE id = $9 RETURNING updated_at`

	err := r.DB.QueryRow(query, news.Title, news.Slug, news.Summary, news.Content, news.ImageURL, news.Status, news.CategoryID, news.AuthorID, news.ID).Scan(&news.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes a news article by ID
func (r *NewsRepository) Delete(id int64) error {
	query := `DELETE FROM news WHERE id = $1`

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
