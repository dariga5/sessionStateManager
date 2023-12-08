package session_pool

import (
	session "SSM/main/session_s"
	"errors"
)

type Pool struct {
}

func (p Pool) AddSession(s *session.SessionExport) error {
	return errors.New("Error")
}

func (p Pool) FindSession(id string) *session.SessionExport {
	return nil
}
