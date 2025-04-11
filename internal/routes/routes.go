package routes

import (
	"net/http"

	"github.com/elltja/news-website/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterRoutes() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	router.Get("/", handlers.HomePageHandler)
	router.Get("/article/{id}", handlers.ArticlePageHandler)

	return router
}
