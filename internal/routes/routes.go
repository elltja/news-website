package routes

import (
	"net/http"

	"github.com/elltja/news-website/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterRoutes() chi.Router {
	// db := database.DB
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		utils.RenderTemplate(w, "web/templates/index.html", nil)
	})
	router.Get("/article/{id}", func(w http.ResponseWriter, r *http.Request) {
		type data struct {
			Id string
		}
		id := chi.URLParam(r, "id")
		utils.RenderTemplate(w, "web/templates/article.html", &data{
			Id: id,
		})
	})
	return router
}
