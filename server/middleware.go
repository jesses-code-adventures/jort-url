package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jesses-code-adventures/jort-url/database"
)

func parseUsernameAndPasswordFromForm(r *http.Request) (string, string, error) {
	return r.FormValue("username"), r.FormValue("password"), nil
}

func parseUsernameAndPasswordFromJson(r *http.Request) (string, string, error) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return "", "", err
	}
	return req.Username, req.Password, nil
}

func (s *Server) parseUsernameAndPassword(r *http.Request) (string, string, error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		return parseUsernameAndPasswordFromJson(r)
	default:
		return parseUsernameAndPasswordFromForm(r)
	}
}

func (s *Server) parseUserDetailsFromCookie(r *http.Request) (int, string, error) {
	cookie, err := r.Cookie("jort_user_id")
	if err != nil {
		return 0, "", fmt.Errorf("No user ID cookie found")
	}
	userId, err := strconv.Atoi(cookie.Value)
	if err != nil {
		return 0, "", fmt.Errorf("Invalid user ID cookie")
	}
	tokenCookie, err := r.Cookie("jort_url_token")
	if err != nil {
		return 0, "", fmt.Errorf("No token cookie found")
	}
	return userId, tokenCookie.Value, nil
}

func (s *Server) userIsAuthenticated(r *http.Request) error {
	userId, token, err := s.parseUserDetailsFromCookie(r)
	if err != nil {
		return err
	}
	err = s.Db.Authenticate(userId, token)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) authenticated(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := s.userIsAuthenticated(r)
		if errors.As(err, &database.InvalidCredentialsError{}) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
