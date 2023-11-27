package session_s

type Session struct {
	id    string
	state sessionState
	cache *sessionCache
}

func createSessionObj() *Session {
	var obj *session = new(Session)

	obj.id = session_utils.GenerateUUID()
	obj.state = on
	obj.cache = new(SessionCache)

	return obj
}
