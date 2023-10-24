package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UUID uuid.UUID

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	Email        string
	HashPassword string
}

type Assignment struct {
	gorm.Model
	Name          string
	Points        int
	NumOfAttempts int
	Deadline      time.Time
}
