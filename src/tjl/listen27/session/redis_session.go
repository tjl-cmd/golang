package session

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"sync"
)

const (
	SessionFlagNone = iota
	SessionFlagModify
	SessionFlagLoad
)

type RedisSession struct {
	sessionId  string
	pool       *redis.Pool
	sessionMap map[string]interface{}
	rwlock     sync.RWMutex
	flag       int
}

func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	s := &RedisSession{
		sessionId:  id,
		sessionMap: make(map[string]interface{}, 8),
		flag:       SessionFlagNone,
		pool:       pool,
	}
	return s
}

func (r *RedisSession) Set(key string, value interface{}) error {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	r.sessionMap[key] = value
	r.flag = SessionFlagModify
}
func (r *RedisSession) loadFromRedis() (err error) {
	conn := r.pool.Get()
	reply, err := conn.Do("GET", r.sessionId)
	if err != nil {
		return
	}
	data, err := redis.String(reply, err)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(data), &r.sessionMap)
	if err != nil {
		return
	}
	return

}
func (r *RedisSession) Get(key string) (result interface{}, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	//实现一个延迟加载的功能
	if r.flag == SessionFlagNone {
		//没有加载,那么就从redis中加载数据
		err = r.loadFromRedis()
		if err != nil {
			return
		}
	}
	result, ok := r.sessionMap[key]
	if !ok {
		err = ErrKeyNotExistnSession
		return
	}
	return
}
func (r *RedisSession) Del(key string) (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	r.flag = SessionFlagModify
	delete(r.sessionMap, key)
	return
}
func (r *RedisSession) Save() (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	if r.flag != SessionFlagModify {
		return
	}
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}
	conn := r.pool.Get()
	_, err = conn.Do("SET", r.sessionId, string(data))
	if err != nil {
		return
	}
	return
}
