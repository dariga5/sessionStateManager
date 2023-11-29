package session_s

import (
	session_utils "SSM/main/session_s/session_utils_lib"
)

type session struct {
	id    string
	state sessionState
	cache *sessionCache
}

type SessionExport session

func CreateSessionObj() *session {
	var obj *session = new(session)

	obj.id = session_utils.GenerateUUID()
	obj.state = on
	obj.cache = new(sessionCache)

	return obj
}

func (session *session) GetID() string {
	return session.id
}
func (session *session) GetState() sessionState {
	return session.state
}
func (session *session) GetCache() *sessionCache {
	return session.cache
}
