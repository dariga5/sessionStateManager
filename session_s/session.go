package session_s

import "encoding/json"

type Session struct {
	Id    int
	Cache map[string]*Data
}

func CreateSessionObj() Session {

	obj := Session{
		Id:    0,
		Cache: make(map[string]*Data),
	}
	return obj
}

func (session Session) PushRequest(req string) {
	session.Cache[req] = CreateDataObj()
}

func (session Session) GetAllData() string {

	jsonString, _ := json.Marshal(session)

	return string(jsonString)
}
