package session_s

type session struct {
	id    string
	state sessionState
	cache *sessionCache
}

func createSessionObj() *session {
	var obj *session = new(session)

	obj.id = ""
	obj.state = on
	obj.cache = new(sessionCache)

	return obj
}
