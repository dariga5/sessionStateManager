package application_utils

import (
	"html/template"
	"net/http"
)

func InitRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("/home", http.HandlerFunc(home))
	router.Handle("/get_hash", http.HandlerFunc(gost_hash))

	return router
}

func home(w http.ResponseWriter, r *http.Request) {
	path := "/home/i-astanin/golang/sessionStateManager/application/utils/html/home.html"

	tmpl, err := template.ParseFiles(path)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func gost_hash(w http.ResponseWriter, r *http.Request) {

}
