package controllers

import (
	"net/http"
	"news/models"
	"news/services"
	"strconv"

	"github.com/labstack/echo/v5"
)

type CreateCategoryRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type CategoryController struct {
	CategoryService *services.CategoryService
}

func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{
		CategoryService: categoryService,
	}
}

// Create creates a new category
//
// @Summary Create a new category
// @Description Creates a new category using its name and slug.
// @Tags Categories
// @Accept json
// @Produce json
// @Param request body CreateCategoryRequest true "Category information"
// @Success 201 {object} models.Category
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /categories [post]
func (ctrl *CategoryController) Create(ctx *echo.Context) error {
	var request CreateCategoryRequest

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
//
// @Summary Get category by ID
// @Description Returns a category using its ID.
// @Tags Categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [get]
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
//
// @Summary Get all categories
// @Description Returns a list of all categories.
// @Tags Categories
// @Produce json
// @Success 200 {array} models.Category
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /categories [get]
func (ctrl *CategoryController) GetAll(ctx *echo.Context) error {
	categories, err := ctrl.CategoryService.GetAll()
	if err != nil {
		return handleError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, categories)
}
