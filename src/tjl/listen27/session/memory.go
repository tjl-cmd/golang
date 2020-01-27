package session

import "sync"

type MemorySession struct {
	data   map[string]interface{}
	id     string
	rwlock sync.RWMutex
}

func NewMemorySession(id string) Session {
	s := &MemorySession{
		id:   id,
		data: make(map[string]interface{}, 8),
	}
	return s
}

func (m *MemorySession) Set(key string, value interface{}) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.data[key] = value
	return
}
func (m *MemorySession) Get(key string) (value interface{}, err error) {
	m.rwlock.RLock()
	defer m.rwlock.Unlock()
	value, ok := m.data[key]
	if !ok {
		err = ErrKeyNotExistnSession
		return
	}
	return
}
func (m *MemorySession) Del(key string) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data, key)
	return
}
func (m *MemorySession) Save() (err error) {
	return
}
