package main

// Routes - Main routing table
func (s *Server) Routes() {
	mw := chainMiddleware(withSession, withLogging)

	oauthRouter := s.Router.PathPrefix("/oauth").Subrouter()
	oauthRouter.Path("/callback").
		Queries("code", "{code}", "state", "{state}").
		HandlerFunc(mw(s.HandleOAuthCallback()))
	oauthRouter.HandleFunc("/authorize", mw(s.HandleRedirectToAuthenticationProvider()))

	s.Router.HandleFunc("/signup", mw(s.HandleSignup())).Methods("GET", "POST")
	s.Router.HandleFunc("/login", mw(s.HandleLogin())).Methods("GET", "POST")
	s.Router.HandleFunc("/", mw(s.HandleHome()))
}
