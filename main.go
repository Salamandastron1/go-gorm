package main

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"type:VARCHAR(30)"`
	LastName  string
	Email     string `gorm:"unique"`
}

func main() {
	db, err := gorm.Open(mysql.Open("root:rootroot@/go_basics"), &gorm.Config{})
	if err != nil {
		panic("Could not connect with Database")
	}

	fmt.Println("Database connection", db)
	db.AutoMigrate(&User{})

	user := createUser("Thony", "Namaste")

	fmt.Println(user)
}

func createUser(first, last string) *User {

	return &User{
		Model: gorm.Model{
			CreatedAt: time.Now(),
		},
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
