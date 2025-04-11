package routes

import (
	"github.com/elltja/news-website/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterRoutes() chi.Router {
	// db := database.DB
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", handlers.HomePageHandler)
	router.Get("/article/{id}", handlers.ArticlePageHandler)
	return router
}
