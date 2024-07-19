package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jesses-code-adventures/jort-url/database"
	"github.com/jesses-code-adventures/jort-url/pages"
)

func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	err := s.userIsAuthenticated(r)
	if err == nil {
		pages.Home().Render(r.Context(), w)
		return
	}
	if errors.As(err, &database.InvalidCredentialsError{}) || err.Error() == "No user ID cookie found" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	} else {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
