// package routes

// import (
// 	"/api/handlers"

// 	"github.com/gorilla/mux"
// )

// func SetHealthzRoute(r *mux.Router) {
// 	r.HandleFunc("/healthz", handlers.HealthzHandler).Methods("GET")
// }

package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/srivardhanreddy01/webapplication_go/api/handlers"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", handlers.HealthzHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
