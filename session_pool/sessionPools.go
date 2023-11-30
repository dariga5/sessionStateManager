package session_pool

import (
	session "SSM/main/session_s"
	"errors"
)

type Pool interface {
	CreatePool() *Pool
	AddSession(s *session.SessionExport) error
	FindSession(id string) *session.SessionExport
}

type DefaultPool struct {
}

func (p DefaultPool) CreatePool() *DefaultPool {
	return new(DefaultPool)
}

func (p DefaultPool) AddSession(s *session.SessionExport) error {

	return errors.New("Error")
}

func (p DefaultPool) FindSession(id string) *session.SessionExport {
	return nil
}
