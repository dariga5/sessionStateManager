package session_s

type session struct {
	id    string
	state sessionState
	cash  *sessionCache
}

func createSessionObj() *session {
	var obj *session = new(session)

	obj.id = generateID()
	obj.state = on
	obj.cash = new(sessionCache)

	return obj
}

// Перенести в отдельный пакет
func generateID() string {
	return "test"
}
