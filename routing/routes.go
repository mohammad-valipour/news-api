package routing

import (
	"log"
	"news/controllers"

	"github.com/labstack/echo/v5"
)

func Register(
	e *echo.Echo,
	userController *controllers.UserController,
	categoryController *controllers.CategoryController,
	newsController *controllers.NewsController,
) {
	api := e.Group("/api/v1")

	// User routes
	users := api.Group("/users")
	users.POST("", userController.Create)
	users.GET("", userController.GetAll)
	users.GET("/:id", userController.GetByID)

	// Category routes
	categories := api.Group("/categories")
	categories.POST("", categoryController.Create)
	categories.GET("", categoryController.GetAll)
	categories.GET("/:id", categoryController.GetByID)

	// News routes
	news := api.Group("/news")
	news.POST("", newsController.Create)
	news.GET("", newsController.GetAll)
	news.GET("/:id", newsController.GetByID)
	news.PUT("/:id", newsController.Update)
	news.DELETE("/:id", newsController.Delete)

	log.Println("API routes registered successfully")
}
