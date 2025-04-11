package handlers

import (
	"fmt"
	"net/http"

	"github.com/elltja/news-website/internal/model"
	"github.com/elltja/news-website/internal/utils"
	"github.com/go-chi/chi/v5"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {

	articles, err := model.GetArticles()

	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}

	data := struct {
		Articles []model.Article
	}{
		Articles: articles,
	}

	utils.RenderTemplate(w, "web/templates/index.html", data)
}

func ArticlePageHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	article, err := model.GetArticleById(id)
	if err != nil {
		fmt.Println("Error: ", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	data := struct {
		Article model.Article
	}{Article: article}

	utils.RenderTemplate(w, "web/templates/article.html", data)
}
