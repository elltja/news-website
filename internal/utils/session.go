package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
)

type sessionData struct {
	UserID string
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

func CreateSession(w http.ResponseWriter, UserID string) error {
	id, err := generateSessionID()
	if err != nil {
		return err
	}
	sessions[id] = sessionData{
		UserID: UserID, // ?
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

func GetServerSession(r *http.Request) (sessionData, error) {
	cookie, err := r.Cookie(CookieSessionKey)
	if err != nil {
		return sessionData{}, errors.New("no session cookie")
	}
	session := sessions[cookie.Value]
	return session, nil
}
