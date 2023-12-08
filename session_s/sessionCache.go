package session_s

import "errors"

const cacheLen = 7

type sessionCache struct {
	lastReqIndex uint8
	lastResIndex uint8

	req [cacheLen]string
	res [cacheLen]string
}

func (sessionCache *sessionCache) GetData(index int) ([2]string, error) {
	var err error

	var req string
	var res string

	if index >= cacheLen || index < 0 {
		err = errors.New("The index goes beyond the boundaries")
	} else {
		err = nil
		req = sessionCache.req[index]
		res = sessionCache.res[index]
	}

	return [2]string{req, res}, err
}

func (sessionCache *sessionCache) PushRequest(data string) (uint8, error) {

	if sessionCache.lastReqIndex >= (cacheLen - 1) {

		sessionCache.lastReqIndex = 0

		return sessionCache.lastReqIndex, errors.New("The cache request was full. Overwriting the queue")
	} else {
		sessionCache.lastReqIndex++

	}

	sessionCache.req[sessionCache.lastReqIndex] = data

	return sessionCache.lastReqIndex, nil
}

func (sessionCache *sessionCache) PushResponse(data string) (uint8, error) {

	if sessionCache.lastResIndex >= (cacheLen - 1) {

		sessionCache.lastResIndex = 0

		return sessionCache.lastResIndex, errors.New("The cache response was full. Overwriting the queue")
	} else {
		sessionCache.lastResIndex++

	}

	sessionCache.res[sessionCache.lastResIndex] = data

	return sessionCache.lastResIndex, nil
}
