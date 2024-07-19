package server

import (
	"encoding/json"
	"net/http"
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
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || (r.Method == "GET" && r.Header.Get("Content-Type") == "www-form-urlencoded") {
		s.createUser(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
