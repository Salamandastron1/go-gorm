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
	Address   Address        `gorm:"foreignKey:UserID"`
}

type Address struct {
	ID     uint
	Name   string
	UserID uint
}

func main() {
	db, err := gorm.Open(mysql.Open("root:rootroot@/go_basics"), &gorm.Config{})
	if err != nil {
		panic("Could not connect with Database")
	}

	fmt.Println("Database connection", db)
	db.Migrator().DropTable(&User{}, &Address{})
	db.Migrator().CreateTable(&User{}, &Address{})

	user := constructUser("Thony", "Namaste")
	fmt.Println(user)

	db.Create(user)
	db.Create(&User{
		Email: sql.NullString{
			String: "fake@email.com",
			Valid:  true,
		},
	})
	u := User{}
	db.Preload("Address").First(&u)
	fmt.Println(u)

}

func constructUser(first, last string) *User {

	return &User{
		FirstName: sql.NullString{String: first, Valid: true},
		LastName:  sql.NullString{String: last, Valid: true},
		Email:     sql.NullString{String: fmt.Sprintf("%s@%s.com", first, last), Valid: true},
		Address: Address{
			Name: "thony street",
		},
	}
}

func createUsers(users []string, db *gorm.DB) {
	for _, v := range users {
		user := strings.Split(v, " ")
		db.Create(constructUser(user[0], user[1]))
	}
}
