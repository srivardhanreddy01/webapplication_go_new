package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/srivardhanreddy01/webapplication_go/cmd/api" // Import your API package
)

type User struct {
	first_name     string
	last_name      string
	email          string
	hashPassword   string
	AccountCreated time.Time
	AccountUpdated time.Time
}

func main() {
	log.Println("Starting the API server...")
	err := http.ListenAndServe(":8081", api.NewRouter())
	if err != nil {
		log.Fatal("Server error:", err)
	}

	csvFile, err := os.Open("/opt/user.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {

		firstname := record[0]
		lastname := record[1]
		email := record[2]
		password := record[3]

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}

		currentTime := time.Now()

		user := User{
			first_name:     firstname,
			last_name:      lastname,
			email:          email,
			hashPassword:   string(hashedPassword),
			AccountCreated: currentTime,
			AccountUpdated: currentTime,
		}

		fmt.Print(user)
	}

}

// package main

// import (
// 	"net/http"
// 	"routes"

// 	"github.com/gorilla/mux"
// )

// func main() {
// 	r := mux.NewRouter()

// 	// Set up healthz route
// 	routes.SetHealthzRoute(r)

// 	// Set up other routes as needed

// 	http.Handle("/", r)
// 	http.ListenAndServe(":8080", nil)
// }
