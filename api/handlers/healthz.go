package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

func isDatabaseHealthy() bool {
	// Define your database connection parameters
	dsn := "root:Sripragna$1@tcp(127.0.0.1:3306)/godatabase"

	// Attempt to open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return false
	}
	defer db.Close()

	// Try to ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging the database:", err)
		return false
	}

	// If no errors occurred, the database connection is healthy
	return true
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")

	// Check database health
	databaseIsHealthy := isDatabaseHealthy()

	if databaseIsHealthy {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}
