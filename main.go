package main

import (
	app_utils "SSM/main/application/utils"
	session_utils "SSM/main/session_s"
	"fmt"
)

const port = ":8054"

func main() {
	router := app_utils.InitRouter()
	sess := session_utils.CreateSessionObj()

	sess.PushRequest("HTTP 1.1 / GET / 123")
	sess.PushRequest("HTTP 1.1 / POST / 242")
	sess.PushRequest("HTTP 1.1 / GET / PARSE")
	sess.PushRequest("HTTP 1.1 / POST / http://info")

	sess.Cache["HTTP 1.1 / GET / 123"].PushResponse("HTTP 1.1 / GET / 321")
	sess.Cache["HTTP 1.1 / GET / 123"].PushResponse("HTTP 1.1 / GET / 987")

	str := sess.GetAllData()

	fmt.Println(str)

	app_utils.StartServer(port, router)
}
