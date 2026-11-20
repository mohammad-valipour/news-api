package services

import (
	"errors"
	"news/models"
	"news/repository"
	"strings"
)

type NewsService struct {
	NewsRepo     *repository.NewsRepository
	CategoryRepo *repository.CategoryRepository
	UserRepo     *repository.UserRepository
}

func NewNewsService(newsRepo *repository.NewsRepository, categoryRepo *repository.CategoryRepository, userRepo *repository.UserRepository) *NewsService {
	return &NewsService{
		NewsRepo:     newsRepo,
		CategoryRepo: categoryRepo,
		UserRepo:     userRepo,
	}
}

// Create creates a new news article.
func (s *NewsService) Create(news *models.News) error {
	news.Title = strings.TrimSpace(news.Title)
	news.Slug = strings.TrimSpace(news.Slug)
	news.Summary = strings.TrimSpace(news.Summary)
	news.Content = strings.TrimSpace(news.Content)

	if news.Title == "" {
		return errors.New("news title is required")
	}

	if news.Slug == "" {
		return errors.New("news slug is required")
	}

	if news.Summary == "" {
		return errors.New("news summary is required")
	}

	if news.Content == "" {
		return errors.New("news content is required")
	}

	if news.Status == "" {
		news.Status = "draft"
	}

	if news.Status != "draft" && news.Status != "published" {
		return errors.New("invalid news status")
	}

	_, err := s.CategoryRepo.GetByID(news.CategoryID)
	if err != nil {
		return errors.New("category not found")
	}

	_, err = s.UserRepo.GetByID(news.AuthorID)
	if err != nil {
		return errors.New("author not found")
	}

	return s.NewsRepo.Create(news)
}

// GetByID returns a news article by ID.
func (s *NewsService) GetByID(id int64) (*models.News, error) {
	if id <= 0 {
		return nil, errors.New("invalid news ID")
	}

	return s.NewsRepo.GetByID(id)
}

// GetAll returns all news articles.
func (s *NewsService) GetAll() ([]models.News, error) {
	return s.NewsRepo.GetAll()
}

// Update updates an existing news article.
func (s *NewsService) Update(news *models.News) error {
	if news.ID <= 0 {
		return errors.New("invalid news ID")
	}

	news.Title = strings.TrimSpace(news.Title)
	news.Slug = strings.TrimSpace(news.Slug)
	news.Summary = strings.TrimSpace(news.Summary)
	news.Content = strings.TrimSpace(news.Content)

	if news.Title == "" {
		return errors.New("news title is required")
	}

	if news.Slug == "" {
		return errors.New("news slug is required")
	}

	if news.Summary == "" {
		return errors.New("news summary is required")
	}

	if news.Content == "" {
		return errors.New("news content is required")
	}

	if news.Status != "draft" && news.Status != "published" {
		return errors.New("invalid news status")
	}

	_, err := s.NewsRepo.GetByID(news.ID)
	if err != nil {
		return errors.New("news not found")
	}

	_, err = s.CategoryRepo.GetByID(news.CategoryID)
	if err != nil {
		return errors.New("category not found")
	}

	_, err = s.UserRepo.GetByID(news.AuthorID)
	if err != nil {
		return errors.New("author not found")
	}

	return s.NewsRepo.Update(news)
}

// Delete deletes a news article by ID.
func (s *NewsService) Delete(id int64) error {
	if id <= 0 {
		return errors.New("invalid news ID")
	}

	return s.NewsRepo.Delete(id)
}
