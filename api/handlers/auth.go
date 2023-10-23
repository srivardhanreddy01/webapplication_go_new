package handlers

import (
	"github.com/srivardhanreddy01/webapplication_go/api/models"
	"gorm.io/gorm"
)

type AuthHandlerDependencies struct {
	DB *gorm.DB
}

func AuthenticateUser(deps AuthHandlerDependencies, email string, password string) (*models.User, error) {
	var user models.User

	// Query the database to retrieve user by email
	result := deps.DB.Where("Email = ?", email).First(&user)

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

// comparePasswords compares the provided password with the stored password hash.
func comparePasswords(storedHash, providedPassword string) bool {
	// Implement your password comparison logic, for example using a library like bcrypt
	// Compare the stored password hash with the provided password
	// If they match, return true; otherwise, return false
	return storedHash == providedPassword
}
