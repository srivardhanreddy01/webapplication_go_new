package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/srivardhanreddy01/webapplication_go/api/handlers"
	"github.com/srivardhanreddy01/webapplication_go/api/middleware"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Define your routes here
	router.HandleFunc("/healthz", handlers.HealthzHandler).Methods("GET")

	router.HandleFunc("/v1/assignments", func(w http.ResponseWriter, r *http.Request) {
		middleware.BasicAuthMiddleware(middleware.AuthMiddlewareDependencies{AuthHandlerDependencies: handlers.AuthHandlerDependencies{DB: db}}, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handlers.GetAllAssignments(w, r, db)
		})).ServeHTTP(w, r)
	}).Methods("GET")

	router.HandleFunc("/v1/assignments", func(w http.ResponseWriter, r *http.Request) {
		middleware.BasicAuthMiddleware(middleware.AuthMiddlewareDependencies{AuthHandlerDependencies: handlers.AuthHandlerDependencies{DB: db}}, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handlers.CreateAssignment(w, r, db)
		})).ServeHTTP(w, r)
	}).Methods("POST")

	// router.HandleFunc("/v1/assignments/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	middleware.BasicAuthMiddleware(middleware.AuthMiddlewareDependencies{AuthHandlerDependencies: handlers.AuthHandlerDependencies{DB: db}}, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		handlers.GetAssignmentById(w, r, db)
	// 	})).ServeHTTP(w, r)
	// }).Methods("GET")

	// router.HandleFunc("/v1/assignments/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println(r.Body)
	// 	fmt.Println(r.Header)
	// 	handlers.GetAssignmentById(w, r, db)
	// }).Methods("GET")

	router.HandleFunc("/v1/assignments/{id}", func(w http.ResponseWriter, r *http.Request) {
		middleware.BasicAuthMiddleware(middleware.AuthMiddlewareDependencies{AuthHandlerDependencies: handlers.AuthHandlerDependencies{DB: db}}, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handlers.GetAssignment(w, r, db)
		})).ServeHTTP(w, r)
	}).Methods("GET")

	router.HandleFunc("/v1/assignments/{id}", func(w http.ResponseWriter, r *http.Request) {
		middleware.BasicAuthMiddleware(middleware.AuthMiddlewareDependencies{AuthHandlerDependencies: handlers.AuthHandlerDependencies{DB: db}}, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handlers.UpdateAssignment(w, r, db)
		})).ServeHTTP(w, r)
	}).Methods("PUT")

	router.HandleFunc("/v1/assignments/{id}", func(w http.ResponseWriter, r *http.Request) {
		middleware.BasicAuthMiddleware(middleware.AuthMiddlewareDependencies{AuthHandlerDependencies: handlers.AuthHandlerDependencies{DB: db}}, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handlers.DeleteAssignment(w, r, db)
		})).ServeHTTP(w, r)
	}).Methods("DELETE")

	return router

}
