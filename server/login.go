package server

import (
	"errors"
	"fmt"
	"github.com/jesses-code-adventures/jort-url/database"
	"github.com/jesses-code-adventures/jort-url/pages"
	"net/http"
)

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	username, password, err := s.parseUsernameAndPassword(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if username == "" || password == "" {
		http.Error(w, "Username and password required", http.StatusBadRequest)
		return
	}
	userId, token, err := s.Db.Login(username, password)
	if errors.As(err, &database.UserNotFoundError{}) {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	} else if errors.As(err, &database.IncorrectPasswordError{}) {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{Name: "jort_url_token", Value: token})
	http.SetCookie(w, &http.Cookie{Name: "jort_user_id", Value: fmt.Sprint(userId)})
	if r.Header.Get("Content-Type") == "application/json" {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(`{"token": "%s", "user_id": %d}`, token, userId)))
		return
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || (r.Method == "GET" && r.Header.Get("Content-Type") == "www-form-urlencoded") {
		s.login(w, r)
	} else if r.Method == "GET" {
		pages.Login().Render(r.Context(), w)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
