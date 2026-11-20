package controllers

import (
	"net/http"
	"news/models"
	"news/services"
	"strconv"

	"github.com/labstack/echo/v5"
)

type NewsController struct {
	NewsService *services.NewsService
}

func NewNewsController(newsService *services.NewsService) *NewsController {
	return &NewsController{
		NewsService: newsService,
	}
}

// Create creates a news article
func (ctrl *NewsController) Create(ctx *echo.Context) error {
	var request struct {
		Title      string `json:"title"`
		Slug       string `json:"slug"`
		Summary    string `json:"summary"`
		Content    string `json:"content"`
		ImageURL   string `json:"image_url"`
		Status     string `json:"status"`
		CategoryID int64  `json:"category_id"`
		AuthorID   int64  `json:"author_id"`
	}

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
		})
	}

	news := &models.News{
		Title:      request.Title,
		Slug:       request.Slug,
		Summary:    request.Summary,
		Content:    request.Content,
		ImageURL:   request.ImageURL,
		Status:     request.Status,
		CategoryID: request.CategoryID,
		AuthorID:   request.AuthorID,
	}

	if err := ctrl.NewsService.Create(news); err != nil {
		return handleError(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, news)
}

// GetByID returns a news article by ID
func (ctrl *NewsController) GetByID(ctx *echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid news id",
		})
	}
	news, err := ctrl.NewsService.GetByID(id)
	if err != nil {
		return handleError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, news)
}

// GetAll returns all news articles
func (ctrl *NewsController) GetAll(ctx *echo.Context) error {
	news, err := ctrl.NewsService.GetAll()
	if err != nil {
		return handleError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, news)
}

// Update updates a news article
func (ctrl *NewsController) Update(ctx *echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid news id",
		})
	}

	var request struct {
		Title      string `json:"title"`
		Slug       string `json:"slug"`
		Summary    string `json:"summary"`
		Content    string `json:"content"`
		ImageURL   string `json:"image_url"`
		Status     string `json:"status"`
		CategoryID int64  `json:"category_id"`
		AuthorID   int64  `json:"author_id"`
	}

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
		})
	}

	news := &models.News{
		ID:         id,
		Title:      request.Title,
		Slug:       request.Slug,
		Summary:    request.Summary,
		Content:    request.Content,
		ImageURL:   request.ImageURL,
		Status:     request.Status,
		CategoryID: request.CategoryID,
		AuthorID:   request.AuthorID,
	}

	if err := ctrl.NewsService.Update(news); err != nil {
		return handleError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, news)
}

// Delete deletes a news article
func (ctrl *NewsController) Delete(ctx *echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid news id",
		})
	}

	if err := ctrl.NewsService.Delete(id); err != nil {
		return handleError(ctx, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}
