package utils

import (
	"net/http"

	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type sessionData struct {
	UserID string
	Role   string
}

const CookieSessionKey = "session-id"

var sessions = make(map[string]sessionData)

func generateSessionID() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func CreateSession(w http.ResponseWriter, UserID string, Role string) error {
	id, err := generateSessionID()
	if err != nil {
		return err
	}
	sessions[id] = sessionData{
		UserID: UserID,
		Role:   Role,
	}
	http.SetCookie(w, &http.Cookie{
		Name:     CookieSessionKey,
		Value:    id,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	})
	return nil
}

func GetSession(r *http.Request) (sessionData, error) {
	cookie, err := r.Cookie(CookieSessionKey)
	if err != nil {
		return sessionData{}, errors.New("no session cookie")
	}
	session := sessions[cookie.Value]
	return session, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func ComparePasswords(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AuthorizeAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionData, err := GetSession(r)
		if err != nil {
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		if sessionData.Role != "admin" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
