package main

import "SSM/main/session_s"

func main() {
	var s = session_s.CreateSessionObj()

	print(s.GetCache().PushRequest("1"))
	print(s.GetCache().PushRequest("2"))
	print(s.GetCache().PushRequest("3"))
	print(s.GetCache().PushRequest("4"))
	print(s.GetCache().PushRequest("5"))
	print(s.GetCache().PushRequest("6"))
	print(s.GetCache().PushRequest("7"))
	print(s.GetCache().PushRequest("8"))
	print(s.GetCache().PushRequest("9"))
	print(s.GetCache().PushRequest("0"))

}
