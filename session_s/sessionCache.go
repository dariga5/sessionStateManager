package session_s

import "errors"

const cacheLen = 7

type sessionCache struct {
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
	var req string
	var res string

	if index > cacheLen || index < cacheLen {
		err = errors.New("Index exceeds cache capacity")
		req = "invalid"
		res = "invalid"

	} else {
		err = nil
		req = sessionCache.req[index]
		res = sessionCache.res[index]
	}

	return cacheExp{
		Err: err,
		Data: [2]string{
			req,
			res,
		},
	}
}

func (sessionCache *sessionCache) PushRequest(data string) error {
	if sessionCache.lastReqIndex == cacheLen {
		return errors.New("Session cache is full")
	} else {
		sessionCache.req[sessionCache.lastReqIndex] = data
		sessionCache.lastReqIndex++
	}
	return nil
}

func (sessionCache *sessionCache) PushResponse(data string) error {
	if sessionCache.lastResIndex == cacheLen {
		return errors.New("Session cache is full")
	} else {
		sessionCache.res[sessionCache.lastResIndex] = data
		sessionCache.lastResIndex++
	}
	return nil
}
