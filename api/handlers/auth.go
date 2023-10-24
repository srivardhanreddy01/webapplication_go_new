package handlers

import (
	"github.com/srivardhanreddy01/webapplication_go/api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandlerDependencies struct {
	DB *gorm.DB
}

func AuthenticateUser(deps AuthHandlerDependencies, email string, password string) (*models.User, error) {
	var user models.User

	// Query the database to retrieve user by email
	result := deps.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// User not found
			return nil, nil
		}
		// Database error
		return nil, result.Error
	}

	// Verify the provided password against the stored password hash
	if !comparePasswords(user.HashPassword, password) {
		// Passwords do not match
		return nil, nil
	}

	// Authentication successful, return the authenticated user
	return &user, nil
}

func comparePasswords(storedHash, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(providedPassword))
	if err != nil {
		return false // Passwords do not match
	}
	return true // Passwords match
}
