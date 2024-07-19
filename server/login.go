package server

import (
	"fmt"
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
	token, err := s.Db.Login(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	http.SetCookie(w, &http.Cookie{Name: "jort_url_token", Value: token})
	w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, token)))
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || (r.Method == "GET" && r.Header.Get("Content-Type") == "www-form-urlencoded") {
		s.login(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
