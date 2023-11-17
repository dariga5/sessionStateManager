package session_s

const cacheLen = 7

type sessionCache struct {
	request  [cacheLen]string
	response [cacheLen]string
}
