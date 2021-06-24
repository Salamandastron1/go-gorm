package users_test

import (
	users "github/salamandastron1/go-basics/Users"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func generateRandomNames() []string {
	var names []string
	for i := 0; i < gofakeit.Number(0, 20); i++ {
		names = append(names, gofakeit.Name())
	}
	return names
}

func TestCreateUsers(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:rootroot@/go_basics"), &gorm.Config{})
	if err != nil {
		panic("Could not connect with Database")
	}

	cases := []struct {
		name  string
		err   bool
		names []string
	}{
		{
			name:  "Users are put into sequentially",
			err:   false,
			names: generateRandomNames(),
		},
		{
			name:  "User is submitted with more than first and last name",
			err:   true,
			names: (append(generateRandomNames(), "Meow mix Meow")),
		},
	}

	for _, v := range cases {
		err := users.CreateUsers(v.names, db)

		if err != nil && v.err == false {
			t.Errorf("Unexpected failure; output is %v", err)
		}

		u := users.User{}
		db.Preload("Book").First(&u)
		uName := u.FirstName.String + " " + u.LastName.String
		if v.names[0] != uName {
			t.Errorf("Database corrupted, non sequential addition found %v", uName)
		}
	}
}
