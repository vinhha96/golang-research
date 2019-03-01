package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Username string `json:"username"`
	Password string `json:"-"`
}

var userList = []User{
	{Username: "user1", Password: "pass1"},
	{Username: "user2", Password: "pass2"},
	{Username: "user3", Password: "pass3"},
}

func IsUserValid(username, password string) bool {
	for _, u := range userList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func RegisterNewUser(username, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("Password can't not empty")
	} else if !isUserNameAvailable(username) {
		return nil, errors.New("Username is not available")
	}

	newUser := User{Username: username, Password: password}
	userList = append(userList, newUser)

	return &newUser, nil
}

func isUserNameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true
}
