package main

import (
	"fmt"
	"log"

	"github.com/elltja/news-website/internal/database"
	"github.com/elltja/news-website/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	err = database.OpenDatabase()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	s := server.NewServer()
	log.Fatal(s.ListenAndServe())
}
