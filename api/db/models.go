package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type Post struct {
	gorm.Model
	OwnerID   uint
	Title     string
	Content   string
	Published bool
}
