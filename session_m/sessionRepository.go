package session_m

import (
	session "SSM/main/session_s"
)

type Repository interface {
	AddSession(s *session.SessionExport) error
	UpdateSession(id string, state session.SessionStateExport) error
	FindSession(id string) *session.SessionExport
}
