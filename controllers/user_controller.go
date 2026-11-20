package controllers

import (
	"net/http"
	"news/models"
	"news/services"
	"strconv"

	"github.com/labstack/echo/v5"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// Create creates a new user
func (ctrl *UserController) Create(ctx *echo.Context) error {
	var request struct {
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Email        string `json:"email"`
		PasswordHash string `json:"password_hash"`
	}

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
		})
	}

	user := &models.User{
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		Email:        request.Email,
		PasswordHash: request.PasswordHash,
		Role:         "author",
	}

	if err := ctrl.UserService.Create(user); err != nil {
		return handleError(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, user)
}

// GetByID returns a user by ID
func (ctrl *UserController) GetByID(ctx *echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid user ID",
		})
	}

	user, err := ctrl.UserService.GetByID(id)
	if err != nil {
		return handleError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, user)

}

// GetAll returns all users
func (ctrl *UserController) GetAll(ctx *echo.Context) error {
	users, err := ctrl.UserService.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, users)
}
