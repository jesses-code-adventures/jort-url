package server

import (
	"encoding/json"
	"fmt"
	"github.com/jesses-code-adventures/jort-url/urls"
	"net/http"
)

func (s *Server) parseUrl(r *http.Request) (string, error) {
	var req struct {
		Url string `json:"url"`
	}
	if r.Header.Get("Content-Type") == "application/json" {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return "", err
		}
	} else {
		req.Url = r.FormValue("url")
	}
	return req.Url, nil
}

func (s *Server) shortenUrl(w http.ResponseWriter, r *http.Request) {
	userId, _, err := s.parseUserDetailsFromCookie(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	url, err := s.parseUrl(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	existing, err := s.Db.GetExistingShortenedPath(userId, url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var shortUrl string
	if existing != nil {
		shortUrl = *existing
	} else {
		newExists := true
		for newExists {
			shortUrl = urls.GetRandomPath()
			newExists, err = s.Db.ShortPathHasBeenUsed(shortUrl)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		err := s.Db.CreateUrl(userId, url, shortUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("%s/%s", r.Host, shortUrl)))
}

func (s *Server) urlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		s.shortenUrl(w, r)
	} else {
		fmt.Println(r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
