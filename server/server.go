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

// Get the static routes that the server serves.
func (s *Server) routes() []string {
	return []string{"static", "user", "login", "logout", "url"}
}

// registerRoutes registers the server routes with optional middleware.
func (s *Server) registerRoutes() {
	staticDir := http.Dir("static")
	fileServer := http.FileServer(staticDir)
	s.Mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	s.Mux.HandleFunc("/", s.rootHandler)
	s.Mux.HandleFunc("/user", s.userHandler)
	s.Mux.HandleFunc("/login", s.loginHandler)
	s.Mux.HandleFunc("/logout", s.authenticated(s.logoutHandler))
	s.Mux.HandleFunc("/url", s.authenticated(s.urlHandler))
	s.Mux.HandleFunc("/{shortenedPath}", s.redirectHandler)
}

func (s *Server) Close() error {
	return s.Db.Close()
}
