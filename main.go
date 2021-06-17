package main

import (
	users "github/salamandastron1/go-basics/Users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:rootroot@/go_basics"), &gorm.Config{})
	if err != nil {
		panic("Could not connect with Database")
	}
	newUsers := []string{
		"Thony Namaste",
		"Angel Ortiz",
		"Nikisha Guadalupe",
		"Randles Carington",
		"Timsey Lohan",
		"Sophia LoveYouLongTime",
		"Karry Unicorn",
	}
	users.CreateUsers(newUsers, db)

}
