package session_pool

import (
	session "SSM/main/session_s"
	"container/list"
	"errors"
)

type Pool interface {
	AddSession(s *session.SessionExport) error
	FindSession(id string) *session.SessionExport
}

type DefaultPool struct {
	list list.List
}

func AddSession(s *session.SessionExport) error {
	var el = list.New().PushFront(s)

	if el != nil {
		return nil
	}
	return errors.New("Error")
}

func FindSession(id string) *session.SessionExport {
	return nil
}
