package session

var (
	sessionMgr SessionMgr
)

func Init(proviter string, addr string, options ...string) (err error) {
	switch proviter {
	case "memory":
		sessionMgr = NewMemorySession()
	}
}
