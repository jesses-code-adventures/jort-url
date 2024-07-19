package server

import (
	"net/http"
)

func (s *Server) redirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	shortenedPath := r.URL.Path[1:]
	url, err := s.Db.GetUrl(shortenedPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if url == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	err = s.Db.IncrementUrlClicks(shortenedPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, *url, http.StatusFound)
}
