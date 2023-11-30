package main

import (
	"SSM/main/session_s"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":7071", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	var s = session_s.CreateSessionObj()

	s.GetCache().PushRequest(r.URL.Path)

	var data = s.GetCache().GetData(7)

	fmt.Println(data.Data)
	fmt.Println(data.Err)
}
