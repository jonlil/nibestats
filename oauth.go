package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jonlil/nibestats/models"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// HandleRedirectToAuthenticationProvider - Method for redirecting page granting access to this application
func (s *Server) HandleRedirectToAuthenticationProvider() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sess, _ := globalSessions.SessionStart(w, r)
		defer sess.SessionRelease(w)

		if sess.Get("UserId") == nil {
			fmt.Println("Please sign in before using this.")
		} else {
			http.Redirect(w, r, fmt.Sprintf("%s/oauth/authorize?response_type=code&client_id=%s&scope=%s&redirect_uri=%s&state=%s",
				s.Nibe.Endpoint,
				s.Nibe.ClientID,
				"READSYSTEM",
				s.Nibe.OAuhRedirectURI,
				"?",
			), 302)
		}
	}
}

// HandleOAuthCallback - Method for receiving authentication token from Nibe
func (s *Server) HandleOAuthCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sess, _ := globalSessions.SessionStart(w, r)
		defer sess.SessionRelease(w)

		params := mux.Vars(r)

		resp, _ := http.PostForm(
			fmt.Sprintf("%s/oauth/token", s.Nibe.Endpoint),
			url.Values{
				"grant_type":    {"authorization_code"},
				"client_id":     {s.Nibe.ClientID},
				"client_secret": {s.Nibe.ClientSecret},
				"code":          {params["code"]},
				"scope":         {"READSYSTEM"},
				"redirect_uri":  {s.Nibe.OAuhRedirectURI},
			})
		defer resp.Body.Close()

		body, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		tokenData := &models.AccessToken{}
		err := json.Unmarshal(body, &tokenData)
		if err != nil {
			fmt.Println("whoops:", err)
		} else {
			user := &models.User{}
			s.DB.First(&user, sess.Get("UserID"))
			tokenData.User = *user
			s.DB.Create(&tokenData)

			fmt.Println("response Status:", resp.Status)
			fmt.Println("response Headers:", resp.Header)
		}
	}
}
