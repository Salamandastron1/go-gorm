package main

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

func main() {
	db, err := gorm.Open(mysql.Open("root:rootroot@/go_basics"), &gorm.Config{})
	if err != nil {
		panic("Could not connect with Database")
	}

	fmt.Println("Database connection", db)
	db.AutoMigrate(&User{})

	users := []string{
		"John Doe",
		"Tim Garrity",
		"Alex Rodriguez",
	}

	createUsers(users, db)

}

func createUser(first, last string) *User {

	return &User{
		FirstName: first,
		LastName:  last,
		Email:     fmt.Sprintf("%s@%s.com", first, last),
	}
}

func createUsers(users []string, db *gorm.DB) {
	for _, v := range users {
		user := strings.Split(v, " ")
		db.Create(createUser(user[0], user[1]))
	}
}
