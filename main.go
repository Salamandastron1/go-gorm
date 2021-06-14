package main

import (
	"fmt"
	"strings"
	"time"

	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName sql.NullString `gorm:"type:VARCHAR(30); null"`
	LastName  sql.NullString `gorm:"size:100; default:'Smith"`
	Email     sql.NullString `gorm:"unique; not null"`
}

func main() {
	db, err := gorm.Open(mysql.Open("root:rootroot@/go_basics"), &gorm.Config{})
	if err != nil {
		panic("Could not connect with Database")
	}

	fmt.Println("Database connection", db)
	db.Migrator().DropTable(&User{})
	db.Migrator().CreateTable(&User{})

	user := createUser("Thony", "Namaste")
	fmt.Println(user)

	db.Create(user)
	db.Create(&User{
		Email: sql.NullString{
			String: "fake@email.com",
			Valid:  true,
		},
	})

}

func createUser(first, last string) *User {

	return &User{
		Model: gorm.Model{
			CreatedAt: time.Now(),
		},
		FirstName: sql.NullString{String: first, Valid: true},
		LastName:  sql.NullString{String: last, Valid: true},
		Email:     sql.NullString{String: fmt.Sprintf("%s@%s.com", first, last), Valid: true},
	}
}

func createUsers(users []string, db *gorm.DB) {
	for _, v := range users {
		user := strings.Split(v, " ")
		db.Create(createUser(user[0], user[1]))
	}
}
