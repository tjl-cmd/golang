package account

import (
	"example.com/m/src/tjl/mercury/session"
)

func InitSession(provider string, addr string, options ...string) (err error) {
	return session.Init(provider, addr, options...)
}
