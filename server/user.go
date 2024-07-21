package server

import (
	"net/http"
)

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	username, password, err := s.parseUsernameAndPassword(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if username == "" || password == "" {
		http.Error(w, "Username and password required", http.StatusBadRequest)
		return
	}
	err = s.Db.CreateUser(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.login(w, r)
}

func (s *Server) userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || (r.Method == "GET" && r.Header.Get("Content-Type") == "www-form-urlencoded") {
		s.createUser(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
