package session_m

import (
	session "SSM/main/session_s"
)

const sz = 4

type memCachePool struct {
	pool *[sz]session.SessionExport
}

func (p *memCachePool) CreatePool() *memCachePool {
	var obj *memCachePool = new(memCachePool)
	obj.pool = new([sz]session.SessionExport)

	return obj
}

func (p *memCachePool) AddSession(s *session.SessionExport) error {
	var err error

	return err
}
func (p *memCachePool) UpdateSession(id string, state session.SessionStateExport) error {
	var err error

	return err
}
func (p *memCachePool) FindSession(id string) *session.SessionExport {
	var s *session.SessionExport

	return s
}
