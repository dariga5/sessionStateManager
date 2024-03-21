package session_s

import (
	"errors"
)

const capacityCache = 7

type Data struct {
	index    int
	response []string
}

func CreateDataObj() *Data {
	obj := Data{
		index:    0,
		response: make([]string, capacityCache),
	}

	return &obj
}

func (data *Data) PushResponse(dt string) (int, error) {

	if data.index < 0 || data.index > capacityCache {
		return data.index, errors.New("The Cache capacity limit has been exceeded")
	}

	data.response[data.index] = dt
	data.index = data.index + 1

	return data.index, nil
}
