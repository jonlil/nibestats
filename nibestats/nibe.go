package nibestats

import (
  "os"
)

const nibeUplinkAPI = "https://api.nibeuplink.com"

// NibeCredentials - struct for accessing credentials
type NibeCredentials struct {
  ClientSecret string
  ClientID string
  OAuhRedirectURI string
  Endpoint string
}

// NewNibeCredentials - Helper for settings common values
func NewNibeCredentials() *NibeCredentials {
  return &NibeCredentials{
    ClientSecret: os.Getenv("CLIENT_SECRET"),
    ClientID: os.Getenv("CLIENT_ID"),
    OAuhRedirectURI: "https://nibe.jl-media.se/oauth/callback",
    Endpoint: nibeUplinkAPI,
  }
}
