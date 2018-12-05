package main

import (
	"fmt"
	"github.com/jonlil/nibestats/models"
	"net/http"
)

// HandleSignup - Handler for signing up
func (s *Server) HandleSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			rnd.HTML(w, http.StatusOK, "signup", nil)
		} else {
			r.ParseForm()

			user := &models.User{
				Email: r.FormValue("email"),
				Name:  r.FormValue("name"),
			}

			user.SetPassword([]byte(r.FormValue("password")))
			s.DB.Create(&user)

			http.Redirect(w, r, "/", 302)
		}
	}
}

func renderBadCredentials(w http.ResponseWriter) {
	rnd.HTML(w, 403, "login", map[string]interface{}{
		"error": "Invalid credentials",
	})
}

// HandleLogin -
func (s *Server) HandleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sess, _ := globalSessions.SessionStart(w, r)
		defer sess.SessionRelease(w)

		if r.Method == "GET" {
			rnd.HTML(w, http.StatusOK, "login", nil)
		} else {
			r.ParseForm()

			user := &models.User{}
			if s.DB.Where("email = ?", r.FormValue("email")).First(&user).RecordNotFound() {
				renderBadCredentials(w)
			} else {
				if user.Authenticate([]byte(r.FormValue("password"))) {
					sess.Set("UserID", user.ID)
					http.Redirect(w, r, "/", 302)
				} else {
					renderBadCredentials(w)
				}
			}
		}
	}
}
