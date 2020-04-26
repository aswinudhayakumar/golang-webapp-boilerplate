package model

import "github.com/jinzhu/gorm"

//Users model
type Users struct {
	gorm.Model

	Name     string
	Email    string
	Password string
}
