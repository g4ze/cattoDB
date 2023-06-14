package router

import (
	"controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Route() {
	router := mux.NewRouter()
	router.HandleFunc("/catto-read-{id}", controller.ReadOne).Methods("GET")
	router.HandleFunc("/catto-read-all", controller.ReadOne).Methods("GET")
	router.HandleFunc("/catto-read-{ids}", controller.ReadOne).Methods("GET")

	// Creating a new server
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	// Start the server
	log.Fatal(server.ListenAndServe())
}
