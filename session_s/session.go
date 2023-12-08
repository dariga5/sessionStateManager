package session_s

import (
	session_utils "SSM/main/session_s/session_utils_lib"
)

type Session struct {
	id    string
	state sessionState
	cache *sessionCache
}

func CreateSessionObj() *Session {
	var obj *Session = new(Session)

	obj.id = session_utils.GenerateUUID()
	obj.state = on
	obj.cache = new(sessionCache)

	return obj
}

func (session *Session) GetID() string {
	return session.id
}
func (session *Session) GetState() sessionState {
	return session.state
}
func (session *Session) GetCache() *sessionCache {
	return session.cache
}
