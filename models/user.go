package models

import (
	"errors"
	"strings"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var userList = []User{
	{Username: "user1", Password: "pass1"},
	{Username: "user2", Password: "pass2"},
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
		return nil, errors.New("the password can't be empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("the username isn't available")
	}

	u := User{Username: username, Password: password}
	userList = append(userList, u)

	return &u, nil
}

func isUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true
}
