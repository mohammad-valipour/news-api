package controllers

import (
	"net/http"
	"news/models"
	"news/services"
	"strconv"

	"github.com/labstack/echo/v5"
)

// CreateUserRequest represents the request body for creating a user.
type CreateUserRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// Create creates a new user
//
// @Summary Create a new user
// @Description Create a new user with the author role
// @Tags Users
// @Accept json
// @Produce json
// @Param Requst body CreateUserRequest true "User Information"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Router /users [post]
func (ctrl *UserController) Create(ctx *echo.Context) error {
	var request CreateUserRequest

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
//
// @Summary Get user by ID
// @Description Returns the details of a user using their ID.
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
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
//
// @Summary Get all users
// @Description Returns a list of all registered users.
// @Tags Users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func (ctrl *UserController) GetAll(ctx *echo.Context) error {
	users, err := ctrl.UserService.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, users)
}
