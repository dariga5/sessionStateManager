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

	_, err1 := sess.Cache["HTTP 1.1 / GET / 123"].PushResponse("987")
	_, err2 := sess.Cache["HTTP 1.1 / GET / 123"].PushResponse("654")
	_, err3 := sess.Cache["HTTP 1.1 / GET / 123"].PushResponse("321")
	_, err4 := sess.Cache["HTTP 1.1 / GET / 123"].PushResponse("987")
	_, err5 := sess.Cache["HTTP 1.1 / GET / 123"].PushResponse("654")
	_, err6 := sess.Cache["HTTP 1.1 / GET / 123"].PushResponse("321")
	_, err7 := sess.Cache["HTTP 1.1 / GET / 123"].PushResponse("987")
	_, err8 := sess.Cache["HTTP 1.1 / GET / 123"].PushResponse("654")
	_, err9 := sess.Cache["HTTP 1.1 / GET / 123"].PushResponse("321")

	fmt.Print(err1)
	fmt.Print(err2)
	fmt.Print(err3)
	fmt.Print(err4)
	fmt.Print(err5)
	fmt.Print(err6)
	fmt.Print(err7)
	fmt.Print(err8)
	fmt.Print(err9)

	app_utils.StartServer(port, router)
}
