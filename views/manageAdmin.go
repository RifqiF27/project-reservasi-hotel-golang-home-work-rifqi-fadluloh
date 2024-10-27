package views

import "fmt"

func ManageAdminReservation() {
	var action int
	fmt.Println("\nAdmin Menu:")
	fmt.Println("1. View All Reservations")
	fmt.Println("2. Confirm a Reservation")
	fmt.Println("3. Cancel a Reservation")
	fmt.Println("4. Update a Reservation")
	fmt.Print("Select an action: ")
	fmt.Scanln(&action)

	switch action {
	case 1:
		ViewAllReservations()
	case 2:
		ConfirmReservation()
	case 3:
		CancelReservationByAdmin()
	case 4:
		UpdateReservation()
	default:
		fmt.Println("Invalid action")
	}
}
