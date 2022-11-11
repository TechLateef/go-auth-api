package models

import (
	"database/sql"
	// "gorm.io/gorm"
	"time"
)

type User struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Email       string `gorm:"unique"`
	Password    string
}
