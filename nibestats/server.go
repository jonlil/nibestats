package nibestats

import (
  "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
  "net/http"
  "log"
)

// Server type
type Server struct {
  Nibe *NibeCredentials
  DB *gorm.DB
  Router *mux.Router
}

// NewServer - initialize server instance
func NewServer() *Server {
  server := &Server{
    Nibe: NewNibeCredentials(),
    Router: mux.NewRouter(),
  }

  return server
}

// Listen - Start listening on http
func (server *Server) Listen() {
  http.Handle("/", server.Router)
  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}
