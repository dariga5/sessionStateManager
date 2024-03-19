package session_s

type sessionState int8

const (
	off   = 0
	on    = 1
	sleep = -1
)

func ChangeState(s *Session, state sessionState) {
	s.state = sessionState(state)
}
