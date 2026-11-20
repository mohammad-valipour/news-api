package main

import (
	"log"
	"net/http"
	"news/config"
	"news/controllers"
	"news/database"
	"news/repository"
	"news/routing"
	"news/services"

	"github.com/labstack/echo/v5"
)

func main() {

	// Load application configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Connect to database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	log.Println("Database connected successfully!")

	// Create repositories
	userRepo := repository.NewUserRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	newsRepo := repository.NewNewsRepository(db)

	// Create services
	userService := services.NewUserService(userRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	newsService := services.NewNewsService(newsRepo, categoryRepo, userRepo)

	// Create controllers
	userController := controllers.NewUserController(userService)
	categoryController := controllers.NewCategoryController(categoryService)
	newsController := controllers.NewNewsController(newsService)

	// Create Echo instance
	e := echo.New()

	// Register routes
	routing.Register(e, userController, categoryController, newsController)

	// Home route
	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Welcome to News Website API",
		})
	})

	// Start server
	log.Println("Starting News Website on port", cfg.AppPort)

	if err := e.Start(":" + cfg.AppPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
