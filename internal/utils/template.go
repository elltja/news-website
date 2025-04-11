package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, path string, data any) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println("Template parsing error")
		handleErr(w, 500)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("Template execution error")
		handleErr(w, 500)
	}
}

func handleErr(w http.ResponseWriter, statusCode int) {
	http.Error(w, "Something went wrong", statusCode)
}
