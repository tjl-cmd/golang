package session

import "myapp/pkg/mod/github.com/pkg/errors@v0.8.1"

var (
	ErrSessionNotExist     = errors.New("session  no exists")
	ErrKeyNotExistnSession = errors.New("key not exists in session")
)
