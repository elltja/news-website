package model

import (
	"database/sql"

	"github.com/elltja/news-website/internal/database"
)

type FullUser struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	Role           string `json:"role"`
}

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserCridentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetFullUserByEmail(email string) (FullUser, error) {
	row := database.DB.QueryRow(`SELECT id, email, hashed_password, role FROM users WHERE email = $1`, email)

	var user FullUser

	if err := row.Scan(&user.ID, &user.Email, &user.HashedPassword, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			return FullUser{}, sql.ErrNoRows
		}
		return FullUser{}, err
	}
	return user, nil
}

func GetUsers() ([]User, error) {
	rows, err := database.DB.Query(`
		SELECT id, email, role FROM users
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(credentials UserCridentials) error {
	_, err := database.DB.Exec(`INSERT INTO users (email, hashed_password) VALUES ($1, $2)`, credentials.Email, credentials.Password)
	if err != nil {
		return err
	}
	return nil
}
