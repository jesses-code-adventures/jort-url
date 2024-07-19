package server

import (
	db "github.com/jesses-code-adventures/jort-url/database"
	"net/http"
)

type Server struct {
	Db  *db.Database
	Mux *http.ServeMux
}

func NewServer() (*Server, error) {
	database, err := db.NewDatabase()
	if err != nil {
		return nil, err
	}
	server := Server{database, http.NewServeMux()}
	server.registerRoutes()
	return &server, nil
}

// registerRoutes registers the server routes with optional middleware.
func (s *Server) registerRoutes() {
	s.Mux.HandleFunc("/user", s.userHandler)
	s.Mux.HandleFunc("/login", s.loginHandler)
	s.Mux.Handle("/logout", s.withMiddleware(http.HandlerFunc(s.logoutHandler), s.authenticated))
}

func (s *Server) Close() error {
	return s.Db.Close()
}
