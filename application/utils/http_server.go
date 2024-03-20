package application_utils

import (
	"log"
	"net/http"
)

func StartServer(port string, router *http.ServeMux) *http.Server {
	server := &http.Server{
		Addr:     port,
		Handler:  router,
		ErrorLog: log.Default(),
	}

	server.ListenAndServe()

	return server
}
