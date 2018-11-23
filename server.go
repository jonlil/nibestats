package nibestats

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"github.com/jonlil/nibe"
)

// Server type
type Server struct {
	Nibe   *nibe.Credentials
	DB     *gorm.DB
	Router *mux.Router
}

// NewServer - initialize server instance
func NewServer() *Server {
	server := &Server{
        Nibe:   nibe.NewCredentials("https://nibe.jl-media.se/oauth/callback"),
		Router: mux.NewRouter(),
	}

	// Install routes
	server.Routes()

	return server
}

// Listen - Start listening on http
func (server *Server) Listen() {
	http.Handle("/", server.Router)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
