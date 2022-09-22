package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique_index;not null" json:"username"`
	Email    string `gorm:"unique_index;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Names    string `json:"names"`
}

type LoginInput struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}
