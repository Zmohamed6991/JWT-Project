package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `json:"full_name"`
	Email    string `json:"email" gorm:"unique" require:"@"`
	Password string `json:"password"`
}

type UserLogin struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}
