// /Users/spulakan/CloudProject/webapplication_go/api/database.go
package api

import (
	"fmt"

	"github.com/srivardhanreddy01/webapplication_go/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {

	dsn := "root:Sripragna$1@tcp(127.0.0.1:3306)/godatabase"
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to the database")
	}

	// Perform initial migrations
	initialMigration()
}

func initialMigration() {
	// Migrate the schema
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to migrate the schema")
	}
}
