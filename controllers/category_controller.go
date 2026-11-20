package controllers

import (
	"net/http"
	"news/models"
	"news/services"
	"strconv"

	"github.com/labstack/echo/v5"
)

type CategoryController struct {
	CategoryService *services.CategoryService
}

func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{
		CategoryService: categoryService,
	}
}

// Create creates a new category
func (ctrl *CategoryController) Create(ctx *echo.Context) error {
	var request struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	}

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
		})
	}

	category := &models.Category{
		Name: request.Name,
		Slug: request.Slug,
	}

	if err := ctrl.CategoryService.Create(category); err != nil {
		return handleError(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, category)
}

// GetByID returns a category by ID
func (ctrl *CategoryController) GetByID(ctx *echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
		})
	}

	category, err := ctrl.CategoryService.GetByID(id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, category)
}

// GetAll returns all categories
func (ctrl *CategoryController) GetAll(ctx *echo.Context) error {
	categories, err := ctrl.CategoryService.GetAll()
	if err != nil {
		return handleError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, categories)
}
