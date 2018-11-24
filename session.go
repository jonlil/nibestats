package main

import (
	"github.com/astaxie/beego/session"
	// Dialect import, not used directly
	_ "github.com/astaxie/beego/session/memcache"
)

var globalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "nibesessionid",
		Gclifetime:      3600,
		EnableSetCookie: true,
		Secure:          false,
		ProviderConfig:  "127.0.0.1:11211",
	}
	globalSessions, _ = session.NewManager("memcache", sessionConfig)
	go globalSessions.GC()
}
