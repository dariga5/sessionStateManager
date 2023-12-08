package session_pool

import (
	session "SSM/main/session_s"
	"errors"
)

const pool_size = 4

type Pool struct {
	pool [pool_size]session.Session
}

func (p Pool) AddSession(s *session.Session) error {

	return errors.New("Session pool is full")
}

func (p Pool) FindSession(id string) (*session.Session, error) {
	var session *session.Session
	var err error

	return session, err
}
