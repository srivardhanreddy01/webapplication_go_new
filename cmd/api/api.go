package api

import (
	"github.com/gorilla/mux"
	"github.com/srivardhanreddy01/webapplication_go/api/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Define your routes here
	router.HandleFunc("/healthz", handlers.HealthzHandler).Methods("GET")

	return router
}
