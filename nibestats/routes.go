package nibestats

// Routes - Main routing table
func (s *Server) Routes() {
  oauthRouter := s.Router.PathPrefix("/oauth").Subrouter()
  oauthRouter.HandleFunc("/authorize", s.HandleRedirectToAuthenticationProvider())
}
