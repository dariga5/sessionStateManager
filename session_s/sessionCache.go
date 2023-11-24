package session_s

const cacheLen = 7

type sessionCache struct {
	req [cacheLen]string
	res [cacheLen]string
}
