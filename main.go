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
	users := []string{
		"John Doe",
		"Tim Garrity",
		"Alex Rodriguez",
	}
	db.AutoMigrate(&User{})

	for _, v := range users {
		user := strings.Split(v, " ")
		db.Create(createUser(user[0], user[1]))
	}

	// db.Create(createUser("John", "Doe"))
	// user := createUser("John 2", "doe 2")
	// user.ID = 2
	// db.Updates(user)
	// db.Delete(&User{ID: 1})
	// db.Delete(&User{ID: 2})
	// db.Delete(&User{ID: 3})

}

func createUser(first, last string) *User {

	return &User{
		FirstName: first,
		LastName:  last,
		Email:     fmt.Sprintf("%s@%s.com", first, last),
	}
}
