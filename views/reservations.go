package views

import (
	"fmt"
	"main/models"
	"main/services"
	"main/utils"
	"strings"
	"time"
)

func ViewAllReservations() {
	reservations := services.ViewAllReservations()
	fmt.Printf("\nAll Reservations:\n")
	fmt.Println(strings.Repeat("-", 40))

	for _, r := range reservations {
		fmt.Printf("ID: %d\nCustomer: %s\nPhone Number: %s\nRoom ID: %d\nReservation Date: %s\nTotal Payment: %d\nPayment Status: %s\nReservation Status: %s\nCreateAt: %s\n",
			r.ID, r.CustomerName, r.CustomerPhone, r.RoomID, r.ReservationDate, r.TotalPayment, r.PaymentStatus, r.ReservationStatus, r.CreatedAt)
		fmt.Println(strings.Repeat("-", 40))
	}
}

func ConfirmReservation() {
	var id int
	id, err := utils.InputInt("Enter Reservation ID to confirm: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = services.ConfirmReservation(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Reservation confirmed successfully!")
	}
}

func CancelReservationByAdmin() {
	var id int
	id, err := utils.InputInt("Enter Reservation ID to cancel: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = services.CancelReservationByAdmin(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Reservation canceled successfully!")
	}
}

func UpdateReservation() {
	var id, roomID, totalPayment int
	id, err := utils.InputInt("Enter Reservation ID to update: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	roomID, err = utils.InputInt("Enter new Room ID: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	totalPayment, err = utils.InputInt("Enter new Total Payment: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = services.UpdateReservation(id, roomID, totalPayment)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Reservation updated successfully!")
	}
}

func MakeReservation() {
	var roomID, totalPayment int
	var reservationDateStr string

	roomID, err := utils.InputInt("Enter Room ID: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var room models.Room
	for _, r := range services.Rooms {
		if r.ID == roomID && r.IsAvailable {
			room = r
			break
		}
	}

	if room.ID == 0 {
		fmt.Println("Room not found or not available.")
		return
	}

	fmt.Println("Facilities included:")
	for _, facility := range room.Facilities {
		fmt.Printf("- %s\n", facility.Name)
	}

	reservationDateStr, err = utils.InputStr("Enter Reservation Date (YYYY-MM-DD): ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	reservationDate, err := time.Parse("2006-01-02", reservationDateStr)
	if err != nil {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
		return
	}

	totalPayment, err = utils.InputInt("Enter Total Payment: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = services.MakeReservation(currentUser.Username, currentUser.PhoneNumber, roomID, reservationDate, totalPayment)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Reservation made successfully!")
	}
}

func CancelReservation() {
	var id int
	id, err := utils.InputInt("Enter Reservation ID to cancel: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = services.CancelReservation(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Reservation canceled successfully!")
	}
}

func ViewMyReservations(customerName string) {
	reservations := services.ViewMyReservations(customerName)
	fmt.Printf("\nMy Reservations:\n")
	fmt.Println(strings.Repeat("-", 40))

	if len(reservations) == 0 {
		fmt.Println("no reservation")
	} else {
		for _, r := range reservations {
			fmt.Printf("ID: %d\nCustomer: %s\nPhone Number: %s\nRoom ID: %d\nReservation Date: %s\nTotal Payment: %d\nPayment Status: %s\nReservation Status: %s\nCreateAt: %s\n",
				r.ID, r.CustomerName, r.CustomerPhone, r.RoomID, r.ReservationDate, r.TotalPayment, r.PaymentStatus, r.ReservationStatus, r.CreatedAt)
			fmt.Println(strings.Repeat("-", 40))
		}
	}
}
