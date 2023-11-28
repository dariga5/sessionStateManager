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

	if index > cacheLen || index < 0 {
		err = errors.New("The index goes beyond the boundaries")
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

func (sessionCache *sessionCache) PushRequest(data string) cacheExp {
	var err error

	if sessionCache.lastReqIndex > cacheLen {
		sessionCache.lastReqIndex = 0
		err = errors.New("The cache request was full. Overwriting the queue")
	} else {
		sessionCache.lastReqIndex++
	}

	sessionCache.req[sessionCache.lastReqIndex] = data

	return cacheExp{
		Err:  err,
		Data: sessionCache.lastReqIndex,
	}
}

func (sessionCache *sessionCache) PushResponse(data string) cacheExp {
	var err error

	if sessionCache.lastResIndex > cacheLen {
		sessionCache.lastResIndex = 0
		err = errors.New("The cache response was full. Overwriting the queue")
	} else {
		sessionCache.lastResIndex++
	}

	sessionCache.res[sessionCache.lastResIndex] = data

	return cacheExp{
		Err:  err,
		Data: sessionCache.lastResIndex,
	}
}
