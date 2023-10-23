package models

import (
	"time"
)

type User struct {
	ID             uint `gorm:"primaryKey"`
	FirstName      string
	LastName       string
	Email          string
	HashPassword   string
	AccountCreated time.Time
	AccountUpdated time.Time
}
