package server

import (
	"net/http"
)

func (s *Server) logout(w http.ResponseWriter, r *http.Request) {
	userId, token, err := s.parseUserDetailsFromCookie(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.Db.Logout(userId, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	http.SetCookie(w, &http.Cookie{Name: "jort_user_id", MaxAge: -1, Value: ""})
	http.SetCookie(w, &http.Cookie{Name: "jort_url_token", MaxAge: -1, Value: ""})
	w.Write([]byte(`{"logged_out": "success"}`))
}

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		s.logout(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
