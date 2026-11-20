package services

import (
	"errors"
	"news/models"
	"news/repository"
	"strings"
)

type CategoryService struct {
	CategoryRepo *repository.CategoryRepository
}

func NewCategoryService(categoryRepo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		CategoryRepo: categoryRepo,
	}
}

// Create creates a new category
func (s *CategoryService) Create(category *models.Category) error {
	category.Name = strings.TrimSpace(category.Name)
	category.Slug = strings.TrimSpace(category.Slug)

	if category.Name == "" {
		return errors.New("category name is required")
	}

	if category.Slug == "" {
		return errors.New("category slug is required")
	}

	return s.CategoryRepo.Create(category)
}

// GetByID returns a category by ID
func (s *CategoryService) GetByID(id int64) (*models.Category, error) {
	if id <= 0 {
		return nil, errors.New("invalid ID")
	}

	return s.CategoryRepo.GetByID(id)
}

// GetAll returns all categories
func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.CategoryRepo.GetAll()
}
