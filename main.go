package main

import (
	"fmt"
	"main/utils"
	"main/views"
)

func main() {
	var choice int
	for {
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Select an option: ")
		fmt.Scanln(&choice)
		utils.ClearScreen()
		switch choice {
		case 1:
			views.Register()
		case 2:
			views.Login()
		case 3:
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
