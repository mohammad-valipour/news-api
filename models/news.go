package models

import "time"

type News struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Summary    string    `json:"summary"`
	Content    string    `json:"content"`
	ImageURL   string    `json:"image_url"`
	Status     string    `json:"status"`
	CategoryID int64     `json:"category_id"`
	AuthorID   int64     `json:"author_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
