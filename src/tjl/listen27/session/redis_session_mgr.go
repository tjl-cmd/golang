package session

import (
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type RedisSessionMgr struct {
	addr   string
	passwd string
	pool   *redis.Pool
	rwlock sync.RWMutex
}

func NewRedisSessionMgr() SessionMgr {
	sr := &RedisSessionMgr{
		sessionMap: make(map[string]session, 1024),
	}

	return sr
}
func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Wait:        false,
	}
}
func (r *RedisSessionMgr) Init(addr string, options ...string) (err error) {
	if len(options) > 0 {
		r.passwd = options[0]
	}
	r.pool = newPool(addr, r.passwd)
	r.addr = addr
	return
}

func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	id := uuid.NewV4()
	if err != nil {
		return
	}
	sessionId := id.String()
	session = NewRedisSession(sessionId, r.pool)
	s.sessionMap[sessionId] = session
	return
}
func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	r.rwlock.RLock()
	defer r.rwlock.RUnlock()
	session, ok := r.sessionMap[sessionId]
	if !ok {
		err = ErrSessionNotExist
		return
	}
	return
}
