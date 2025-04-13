package server

import (
	"net/http"

	"github.com/elltja/news-website/internal/server/routes"
	"github.com/elltja/news-website/internal/utils"
)

func NewServer() *http.Server {
	server := &http.Server{
		Addr:    ":" + utils.GetEnvOrDefault("PORT", "8080"),
		Handler: routes.RegisterRoutes(),
	}
	return server
}
