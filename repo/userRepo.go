package repo

import "github.com/jinzhu/gorm"

//SayHello is a function to interact with DB
func SayHello(db *gorm.DB) string {
	//your database query goes here
	return "Hello gophers"
}
