package server

import (
	"errors"
	"net/http"

	"github.com/jesses-code-adventures/jort-url/database"
	"github.com/jesses-code-adventures/jort-url/pages"
)

func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	err := s.userIsAuthenticated(r)
	if err != nil {
		if errors.As(err, &database.InvalidCredentialsError{}) || err.Error() == "No user ID cookie found" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pages.Home().Render(r.Context(), w)
}
