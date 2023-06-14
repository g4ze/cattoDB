package router

import (
	"controller"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Route() {
	router := mux.NewRouter()
	router.HandleFunc("/catto-read-{id}", controller.ReadOne).Methods("GET")
	router.HandleFunc("/catto-read-all", controller.ReadOne).Methods("GET")
	router.HandleFunc("/catto-read-{ids}", controller.ReadOne).Methods("GET")

	// Creating a new server
	server := &http.Server{
		Addr:         "localhost:" + port,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 6 * time.Second,

		IdleTimeout:       15 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           router,
	}

	// Start the server
	log.Fatal(server.ListenAndServe())
}
