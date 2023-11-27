package session_s

import "errors"

const cacheLen = 7

type sessionCache struct {
	cacheLen     uint8
	lastReqIndex uint8
	lastResIndex uint8

	req [cacheLen]string
	res [cacheLen]string
}
type cacheExp struct {
	Err  error
	Data any
}

func (sessionCache *sessionCache) GetData(index int) cacheExp {
	var err error
	if index > cacheLen || index < cacheLen {
		err = errors.New("Index exceeds cache capacity")
	}
	return cacheExp{
		Err: err,
		Data: [2]string{
			sessionCache.req[index],
			sessionCache.res[index],
		},
	}
}

func (sessionCache *sessionCache) SetRequest(data string) error {
	if sessionCache.lastReqIndex == cacheLen {
		return errors.New("Session cache is full")
	} else {
		sessionCache.req[sessionCache.lastReqIndex] = data
		sessionCache.lastReqIndex++
	}
	return nil
}

func (sessionCache *sessionCache) SetResponse(data string) error {
	if sessionCache.lastResIndex == cacheLen {
		return errors.New("Session cache is full")
	} else {
		sessionCache.res[sessionCache.lastResIndex] = data
		sessionCache.lastResIndex++
	}
	return nil
}
