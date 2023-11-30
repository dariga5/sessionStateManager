package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":7071", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

}
