package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenDatabase() error {
	var (
		database = os.Getenv("DB_DATABASE")
		password = os.Getenv("DB_PASSWORD")
		username = os.Getenv("DB_USERNAME")
		port     = os.Getenv("DB_PORT")
		host     = os.Getenv("DB_HOST")
	)

	if database == "" || password == "" || username == "" || port == "" || host == "" {
		return fmt.Errorf("missing one or more required environment variables")
	}
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("invalid DB_PORT value: %v", err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, portInt, username, password, database)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	errPing := DB.Ping()
	if errPing != nil {
		return fmt.Errorf("error pinging database: %v", errPing)
	}

	fmt.Println("Successfully connected!")
	return nil
}

func CloseDatabase() error {
	if DB != nil {
		return DB.Close()
	}
	return fmt.Errorf("database is already closed or not initialized")
}
