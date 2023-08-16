package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Age      int
	Gender   string
	Mobile   string
	Email    string
	Password string
}
