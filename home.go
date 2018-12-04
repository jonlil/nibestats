package main

import (
	"github.com/jonlil/nibestats/models"
	"net/http"
)

// HandleHome -
func (s *Server) HandleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sess, _ := globalSessions.SessionStart(w, r)
		defer sess.SessionRelease(w)
		var user *models.User = nil
		var accessToken *models.AccessToken = nil

		if sess.Get("UserID") != nil {
			user = &models.User{}
			accessToken = &models.AccessToken{}

			if s.DB.First(&user, sess.Get("UserID")).RecordNotFound() {
				user = nil
			}

			if s.DB.Where("user_id = ? AND user_id IS NOT NULL", user.ID).First(&accessToken).RecordNotFound() {
				accessToken = nil
			}
		}

		rnd.HTML(w, http.StatusOK, "home", map[string]interface{}{
			"user":            user,
			"nibe_connection": accessToken,
		})
	}
}
