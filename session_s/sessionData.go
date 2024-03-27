package session_s

import "errors"

const capacityCache = 7

type Data struct {
	index    int
	response []string
}

func CreateDataObj() *Data {
	obj := new(Data)

	obj.index = 0
	obj.response = make([]string, capacityCache)

	return obj
}

func (data *Data) PushResponse(dt string) (int, error) {
	if data.index < 0 || data.index > capacityCache-1 {
		return data.index, errors.New("")
	}
	data.response[data.index] = dt
	data.index++

	return data.index, nil
}
