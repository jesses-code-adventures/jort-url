package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jesses-code-adventures/jort-url/database"
)

func (s *Server) parseUsernameAndPassword(r *http.Request) (string, string, error) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if r.Header.Get("Content-Type") == "application/json" {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return "", "", err
		}
	} else {
		req.Username = r.FormValue("username")
		req.Password = r.FormValue("password")
	}
	return req.Username, req.Password, nil
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

func (s *Server) withMiddleware(handler http.Handler, middleware func(http.Handler) http.Handler) http.Handler {
	return middleware(handler)
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

func (s *Server) authenticated(next http.Handler) http.Handler {
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
