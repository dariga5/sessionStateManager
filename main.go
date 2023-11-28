package main

import (
	"SSM/main/session_s"
	"fmt"
)

func main() {
	var s = session_s.CreateSessionObj()

	var data = (s.GetCache().PushRequest("1").Err)
	var data1 = (s.GetCache().PushRequest("1").Err)
	var data2 = (s.GetCache().PushRequest("1").Err)
	var data3 = (s.GetCache().PushRequest("1").Err)
	var data4 = (s.GetCache().PushRequest("1").Err)
	var data5 = (s.GetCache().PushRequest("1").Err)
	var data6 = (s.GetCache().PushRequest("1").Err)
	var data7 = (s.GetCache().PushRequest("1").Err)
	var data8 = (s.GetCache().PushRequest("1").Err)
	var data9 = (s.GetCache().PushRequest("1").Err)

	fmt.Println("Data : ", data)
	fmt.Println("Data : ", data1)
	fmt.Println("Data : ", data2)
	fmt.Println("Data : ", data3)
	fmt.Println("Data : ", data4)
	fmt.Println("Data : ", data5)
	fmt.Println("Data : ", data6)
	fmt.Println("Data : ", data7)
	fmt.Println("Data : ", data8)
	fmt.Println("Data : ", data9)

}
