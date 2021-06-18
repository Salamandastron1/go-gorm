package users_test

import (
	users "github/salamandastron1/go-basics/Users"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func createUsers() []string {
	return []string{
		"Tim Garrity",
		"Thony Namaste",
		"Angel Ortiz",
		"Nikisha Guadalupe",
		"Randles Carington",
		"Timsey Lohan",
		"Sophia LoveYouLongTime",
		"Karry Unicorn",
	}
}

func TestCreateUsers(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:rootroot@/go_basics"), &gorm.Config{})
	if err != nil {
		panic("Could not connect with Database")
	}
	cases := []struct {
		name     string
		expected string
		err      bool
		users    []string
	}{
		{
			name:     "Users are put into sequentially",
			expected: "Tim Garrity",
			err:      false,
			users:    createUsers(),
		},
		{
			name:     "User is submitted with more than first and last name",
			expected: "Tim Garrity",
			err:      true,
			users:    createUsers(),
		},
	}

	for _, v := range cases {
		err := users.CreateUsers(v.users, db)

		if err != nil && v.err == false {
			t.Errorf("Unexpected failure; output is %v", err)
		}
	}
}
