package main

import (
	"net/http"
	"fmt"
)

// HandleHome -
func (s *Server) HandleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sess, _ := globalSessions.SessionStart(w, r)
		defer sess.SessionRelease(w)
		rnd.HTML(w, http.StatusOK, "home", nil)
	}
}
