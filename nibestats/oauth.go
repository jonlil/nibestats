package nibestats

import (
	"fmt"
	"net/http"
)

// HandleRedirectToAuthenticationProvider - Method for redirecting page granting access to this application
func (s *Server) HandleRedirectToAuthenticationProvider() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, fmt.Sprintf("%s/oauth/authorize?response_type=code&client_id=%s&scope=%s&redirect_uri=%s&state=%s",
			s.Nibe.Endpoint,
			s.Nibe.ClientID,
			"READSYSTEM",
			s.Nibe.OAuhRedirectURI,
			"?",
		), 302)
	}
}

