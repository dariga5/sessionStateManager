package main

import (
	"SSM/main/session_pool"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	http.ListenAndServe(":7071", nil)
	var pool = session_pool.DefaultPool.CreatePool()
}
