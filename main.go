package main

import "SSM/main/session_s"

func main() {
	var s = session_s.CreateSessionObj()

	print(s.GetCache().PushRequest("1").Err)
	print(s.GetCache().PushRequest("2").Err)
	print(s.GetCache().PushRequest("3").Err)
	print(s.GetCache().PushRequest("4").Err)
	print(s.GetCache().PushRequest("5").Err)
	print(s.GetCache().PushRequest("6").Err)
	print(s.GetCache().PushRequest("7").Err)
	print(s.GetCache().PushRequest("8").Err)
	print(s.GetCache().PushRequest("9").Err)
	print(s.GetCache().PushRequest("0").Err)

}
