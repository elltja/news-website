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

func GetUserByEmail(email string) (User, error) {
	row := database.DB.QueryRow(`SELECT id, email, hashed_password FROM users WHERE email = $1`, email)

	var user User

	if err := row.Scan(&user.ID, &user.Email, &user.HashedPassword); err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}
	return user, nil
}
