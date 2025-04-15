package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elltja/news-website/internal/model"
	"github.com/elltja/news-website/internal/utils"
	"github.com/go-chi/chi/v5"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {

	articles, err := model.GetArticles()

	if err != nil {
		fmt.Println("Error: ", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
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

func AuthPageHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "web/templates/auth.html", nil)
}

func AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "web/templates/admin.html", nil)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	credentials := model.UserCridentials{}
	err := d.Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := model.GetUserByEmail(credentials.Email)

	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err == sql.ErrNoRows {
		registerUser(credentials)
	} else {
		if !utils.ComparePasswords(credentials.Password, user.HashedPassword) {
			fmt.Println(user.HashedPassword)
			fmt.Println(credentials.Password)
			http.Error(w, "Invalid Password", http.StatusUnauthorized)
			return
		}
	}

	utils.CreateSession(w, user.ID, user.Role)
	w.Write([]byte("Succesfully logged in"))
}

func registerUser(credentials model.UserCridentials) error {
	hashedPassword, _ := utils.HashPassword(credentials.Password)
	err := model.CreateUser(model.UserCridentials{
		Password: hashedPassword,
		Email:    credentials.Email,
	})
	if err != nil {
		return err
	}
	return nil
}
