package main

import (
	"github.com/astaxie/beego/session"
	// Dialect import, not used directly
	_ "github.com/astaxie/beego/session/memcache"
	"os"
)

var globalSessions *session.Manager

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "nibesessionid",
		Gclifetime:      3600,
		EnableSetCookie: true,
		Secure:          false,
		ProviderConfig:  getEnv("MEMCACHED_SESSION_URL", "127.0.0.1:11211"),
	}
	globalSessions, _ = session.NewManager("memcache", sessionConfig)
	go globalSessions.GC()
}
