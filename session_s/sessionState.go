package session_s

type sessionState int8

const (
	on    = 1
	off   = 0
	sleep = 2
)

func changeState(s *session, state sessionState) {
	s.state = sessionState(state)
}
