package model

import (
	"database/sql"

	"github.com/elltja/news-website/internal/database"
)

type User struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

type UserCridentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUserByEmail(email string) (User, error) {
	row := database.DB.QueryRow(`SELECT id, email, hashed_password FROM users WHERE email = $1`, email)

	var user User

	if err := row.Scan(&user.ID, &user.Email, &user.HashedPassword); err != nil {
		if err == sql.ErrNoRows {
			return User{}, sql.ErrNoRows
		}
		return User{}, err
	}
	return user, nil
}

func CreateUser(credentials UserCridentials) error {
	_, err := database.DB.Exec(`INSERT INTO users (email, hashed_password) VALUES ($1, $2)`, credentials.Email, credentials.Password)
	if err != nil {
		return err
	}
	return nil
}
