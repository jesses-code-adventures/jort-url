package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func (s *Server) authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, token, err := s.parseUserDetailsFromCookie(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		err = s.Db.Authenticate(userId, token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
