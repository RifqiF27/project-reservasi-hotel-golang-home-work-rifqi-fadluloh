package views

import (
	"fmt"
	"main/models"
	"main/services"
	"main/utils"
)

var currentUser *models.User

func Register() {
	var username, password, phoneNumber string
	username, err := utils.InputStr("Enter Username: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	password, err = utils.InputStr("Enter Password: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	phoneNumber, err = utils.InputStr("Enter Phone Number: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = services.RegisterUser(username, password, phoneNumber, "customer")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Registration successful!")
	}
}

func Login() {
	var username, password string
	username, err := utils.InputStr("Enter Username: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	password, err = utils.InputStr("Enter Password: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	user, err := services.Login(username, password)
	if err != nil {
		fmt.Println(err)
		return
	}

	currentUser = user
	fmt.Printf("Welcome, %s!\n", user.Username)

	if user.Role == "admin" {
		ManageAdminMenu()
	} else {
		ManageCustomerMenu()
	}
}

func Logout() {
	currentUser = nil
	fmt.Println("Logged out successfully!")
}

func ManageAdminMenu() {

	for {
		fmt.Println("1. Manage Reservation")
		fmt.Println("2. Manage Room")
		fmt.Println("3. View All Customer")
		fmt.Println("4. Logout")
		var option int
		fmt.Scanln(&option)
		utils.ClearScreen()
		switch option {
		case 1:
			ManageAdminReservation()
		case 2:
			ManageRooms()
		case 3:
			ViewAllCustomer()
		case 4:
			fmt.Println("Logging out...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func ManageCustomerMenu() {
	for {
		var action int
		fmt.Println("\nCustomer Menu:")
		fmt.Println("1. Make a Reservation")
		fmt.Println("2. View My Reservations")
		fmt.Println("3. Cancel a Reservation")
		fmt.Println("4. View Available Rooms")
		fmt.Println("5. Logout")
		fmt.Print("Select an action: ")
		fmt.Scanln(&action)
		utils.ClearScreen()
		switch action {
		case 1:
			MakeReservation()
		case 2:
			ViewMyReservations(currentUser.Username)
		case 3:
			CancelReservation()
		case 4:
			ViewAvailableRooms()
		case 5:
			fmt.Println("Logging out...")
			return
		default:
			fmt.Println("Invalid action")
		}
	}
}
