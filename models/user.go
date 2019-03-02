package models

import (
	"errors"
	"github.com/jinzhu/gorm"
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

type UserStore struct {
	DB *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{DB: db}
}

func (store *UserStore) IsUserValid(username, password string) bool {
	var userList []User

	store.DB.Find(&userList)

	for _, u := range userList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func (store *UserStore) RegisterNewUser(username, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("Password can't not empty")
	} else if !store.isUserNameAvailable(username) {
		return nil, errors.New("Username is not available")
	}

	newUser := User{Username: username, Password: password}
	store.DB.Save(&newUser)

	return &newUser, nil
}

func (store *UserStore) isUserNameAvailable(username string) bool {
	var userList []User

	store.DB.Find(&userList)

	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true
}
