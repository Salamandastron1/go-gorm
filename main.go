package main

import (
	"fmt"
	"strings"

	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint
	FirstName sql.NullString `gorm:"type:VARCHAR(30); null"`
	LastName  sql.NullString `gorm:"size:100; default:'Smith"`
	Email     sql.NullString `gorm:"unique; not null"`
	Book      []Book         `gorm:"many2many:user_books"`
}

type Book struct {
	ID    uint
	Title sql.NullString
}

func main() {
	db, err := gorm.Open(mysql.Open("root:rootroot@/go_basics"), &gorm.Config{})
	if err != nil {
		panic("Could not connect with Database")
	}
	db.Migrator().DropTable(&User{}, &Book{})
	db.AutoMigrate(&User{}, &Book{})

	user := constructUser("Thony", "Namaste")
	db.Create(user)
	fmt.Println(user)
	u := User{}
	db.Preload("Book").First(&u)

	fmt.Println("end", &u)

}

func constructUser(first, last string) *User {

	return &User{
		FirstName: sql.NullString{String: first, Valid: true},
		LastName:  sql.NullString{String: last, Valid: true},
		Email:     sql.NullString{String: fmt.Sprintf("%s@%s.com", first, last), Valid: true},
		Book:      constructBooks([]string{"Eragon", "Warriors", "War and Peace", "Accelerate"}),
	}
}

func constructBooks(titles []string) []Book {
	books := []Book{}
	for _, v := range titles {
		books = append(books, Book{Title: sql.NullString{String: v, Valid: true}})
	}

	return books
}

func createUsers(users []string, db *gorm.DB) {
	for _, v := range users {
		user := strings.Split(v, " ")
		db.Create(constructUser(user[0], user[1]))
	}
}
