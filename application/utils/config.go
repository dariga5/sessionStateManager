package application_utils

import (
	session_m "SSM/main/session_m"
)

type config struct {
	cachecapacity uint8
	sessionRepo   session_m.Repository
}
