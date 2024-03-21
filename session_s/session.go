package session_s

type Session struct {
	Id    int
	Cache map[string]Data
}

func CreateSessionObj() Session {

	obj := Session{
		Id:    0,
		Cache: make(map[string]Data),
	}
	return obj
}

func (session Session) PushRequest(req string) {
	session.Cache[req] = CreateDataObj()
}
