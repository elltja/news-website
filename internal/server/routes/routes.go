package routes

import (
	"net/http"

	"github.com/elltja/news-website/internal/server/handlers"
	"github.com/elltja/news-website/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterRoutes() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	router.Get("/", handlers.HomePageHandler)
	router.Get("/article/{id}", handlers.ArticlePageHandler)
	router.Get("/auth", handlers.AuthPageHandler)
	router.Route("/admin", func(r chi.Router) {
		r.Use(utils.AuthorizeAdmin)
		r.Get("/", handlers.AdminPageHandler)
	})
	router.Route("/api", func(r chi.Router) {
		r.Post("/authenticate", handlers.AuthHandler)
	})

	return router
}
