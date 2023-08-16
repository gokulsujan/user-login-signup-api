package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"Name"`
	Age      int    `json:"Age"`
	Gender   string `json:"Gender"`
	Mobile   string `json:"Mobile"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
