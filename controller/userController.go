package controller

import (
	"github.com/aswinudhayakumar/golang-boilerplate/model"
	"github.com/aswinudhayakumar/golang-boilerplate/repo"
	"github.com/jinzhu/gorm"
)

//UserController is a struct for database
type UserController struct {
	DB *gorm.DB
}

//SayHello is a function to add new user
func SayHello(db *gorm.DB) model.Output {

	var out model.Output
	out.Data = repo.SayHello(db)
	out.Code = 200
	out.Error = nil

	return out

}
