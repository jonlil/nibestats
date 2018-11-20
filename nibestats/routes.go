package nibestats

// Routes - Main routing table
func (s *Server) Routes() {
  oauthRouter := s.Router.PathPrefix("/oauth").Subrouter()
  oauthRouter.Path("/callback").
    Queries("code", "{code}", "state", "{state}").
    HandlerFunc(s.HandleOAuthCallback())
  oauthRouter.HandleFunc("/authorize", s.HandleRedirectToAuthenticationProvider())
}
