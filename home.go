package main

import (
	"net/http"
)

// HandleHome -
func (s *Server) HandleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rnd.HTML(w, http.StatusOK, "home", nil)
	}
}
