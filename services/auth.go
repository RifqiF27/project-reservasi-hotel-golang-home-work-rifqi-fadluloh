package services

import (
	"fmt"
	"main/models"
	"main/utils"
	"os"
)

var users []models.User

func init() {
	err := utils.ReadJSON("data/users.json", &users)
	if err != nil {
		if os.IsNotExist(err) {
			users = []models.User{} 
		} else {
			panic("Could not load users data.")
		}
	}
}

func RegisterUser(username, password, phoneNumber, role string) error {
	for _, user := range users {
		if user.Username == username {
			return fmt.Errorf("username already taken")
		}
	}

	newUser := models.User{
		ID:          len(users) + 1,
		Username:    username,
		Password:    password,
		PhoneNumber: phoneNumber,
		Role:        role,
	}
	users = append(users, newUser)
	utils.WriteJSON("data/users.json", users)

	return nil
}

func Login(username, password string) (*models.User, error) {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("invalid username or password")
}
