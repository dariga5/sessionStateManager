package main

import (
	app_utils "SSM/main/application/utils"
	session_utils "SSM/main/session_s"
	"fmt"
)

const port = ":8054"
const req = "HTTP 1.1 / GET / 123"

func main() {
	router := app_utils.InitRouter()

	sess := session_utils.CreateSessionObj()
	sess.Cache[req] = session_utils.CreateCache()
	sess.Cache[req].PushResponse("1234")
	sess.Cache[req].PushResponse("1234")
	sess.Cache[req].PushResponse("1234")
	sess.Cache[req].PushResponse("1234")
	sess.Cache[req].PushResponse("1234")
	sess.Cache[req].PushResponse("1234")
	sess.Cache[req].PushResponse("1234")

	fmt.Println(sess.GetAllData())

	app_utils.StartServer(port, router)
}
