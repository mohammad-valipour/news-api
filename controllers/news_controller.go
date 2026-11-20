package controllers

import (
	"net/http"
	"news/models"
	"news/services"
	"strconv"

	"github.com/labstack/echo/v5"
)

type CreateNewsRequest struct {
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Summary    string `json:"summary"`
	Content    string `json:"content"`
	ImageURL   string `json:"image_url"`
	Status     string `json:"status"`
	CategoryID int64  `json:"category_id"`
	AuthorID   int64  `json:"author_id"`
}

type UpdateNewsRequest struct {
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Summary    string `json:"summary"`
	Content    string `json:"content"`
	ImageURL   string `json:"image_url"`
	Status     string `json:"status"`
	CategoryID int64  `json:"category_id"`
	AuthorID   int64  `json:"author_id"`
}

type NewsController struct {
	NewsService *services.NewsService
}

func NewNewsController(newsService *services.NewsService) *NewsController {
	return &NewsController{
		NewsService: newsService,
	}
}

// Create creates a news article
//
// @Summary Create a news article
// @Description Creates a new news article.
// @Tags News
// @Accept json
// @Produce json
// @Param request body CreateNewsRequest true "News information"
// @Success 201 {object} models.News
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /news [post]
func (ctrl *NewsController) Create(ctx *echo.Context) error {
	var request CreateNewsRequest

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
//
// @Summary Get news by ID
// @Description Returns a news article using its ID.
// @Tags News
// @Produce json
// @Param id path int true "News ID"
// @Success 200 {object} models.News
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /news/{id} [get]
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
//
// @Summary Get all news
// @Description Returns a list of all news articles.
// @Tags News
// @Produce json
// @Success 200 {array} models.News
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /news [get]
func (ctrl *NewsController) GetAll(ctx *echo.Context) error {
	news, err := ctrl.NewsService.GetAll()
	if err != nil {
		return handleError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, news)
}

// Update updates a news article
//
// @Summary Update a news article
// @Description Updates an existing news article using its ID.
// @Tags News
// @Accept json
// @Produce json
// @Param id path int true "News ID"
// @Param request body UpdateNewsRequest true "Updated news information"
// @Success 200 {object} models.News
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /news/{id} [put]
func (ctrl *NewsController) Update(ctx *echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid news id",
		})
	}

	var request UpdateNewsRequest

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
//
// @Summary Delete a news article
// @Description Deletes a news article using its ID.
// @Tags News
// @Param id path int true "News ID"
// @Success 204 "News deleted successfully"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /news/{id} [delete]
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
