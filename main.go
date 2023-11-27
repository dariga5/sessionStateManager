package main

import (
	"SSM/main/session_s"
)

func main() {

	var s = session_s.CreateSessionObj()

	s.GetCache().SetRequest("1")
	s.GetCache().SetResponse("2")

	s.GetCache().SetRequest("3")
	s.GetCache().SetResponse("4")

	s.GetCache().SetRequest("5")
	s.GetCache().SetResponse("6")

	s.GetCache().SetRequest("7")
	s.GetCache().SetResponse("8")

	s.GetCache().SetRequest("9")
	s.GetCache().SetResponse("10")

	s.GetCache().SetRequest("11")
	s.GetCache().SetResponse("12")

	s.GetCache().SetRequest("13")
	s.GetCache().SetResponse("14")

	s.GetCache().SetRequest("15")
	s.GetCache().SetResponse("16")

	s.GetCache().SetRequest("17")
	s.GetCache().SetResponse("18")

	s.GetCache().SetRequest("19")
	s.GetCache().SetResponse("20")
}
