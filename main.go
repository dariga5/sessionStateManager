package main

import (
	"SSM/main/session_s"
)

func main() {

	var s = session_s.CreateSessionObj()

	s.GetCache().SetRequest("123")
	s.GetCache().SetResponse("234")

	var err = s.GetCache().GetData(0).Err

	print(err.Error())
}
