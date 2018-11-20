package main

import (
  "github.com/gorilla/mux"
  "log"
  "net/http"
  "./nibestats"
)


// ?code={code}&state={state}
func main() {
  rtr := mux.NewRouter()

  // oauth routes
  oauthRouter := rtr.PathPrefix("/oauth").Subrouter()
  oauthRouter.Path("/callback").
      Queries("code", "{code}", "state", "{state}").
      HandlerFunc(nibestats.OAuthCallbackHandler)
  oauthRouter.HandleFunc("/authorize", nibestats.RedirectToAuthenticationProviderHandler)

  http.Handle("/", rtr)

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}
