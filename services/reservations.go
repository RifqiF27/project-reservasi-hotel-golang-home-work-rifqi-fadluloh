package services

import (
	"fmt"
	"main/models"
	"main/utils"
	"time"
)

var reservations []models.Reservation

func init() {
	err := utils.ReadJSON("data/reservations.json", &reservations)
	if err != nil {
		panic("Could not load reservations data.")
	}
}

func ViewAllReservations() []models.Reservation {
	return reservations
}



func ConfirmReservation(id int) error {
	for i, reservation := range reservations {
		if reservation.ID == id {
			if reservation.ReservationStatus == "Pending Approval" {
				reservations[i].ReservationStatus = "Confirmed"
				reservations[i].PaymentStatus = "Paid"
				utils.WriteJSON("data/reservations.json", reservations)
				return nil
			} else {
				return fmt.Errorf("Reservation is already confirmed or canceled.")
			}
		}
	}
	return fmt.Errorf("Error: Reservation not found.")
}

func CancelReservationByAdmin(id int) error {
	for i, reservation := range reservations {
		if reservation.ID == id {
			reservations[i].ReservationStatus = "Canceled"
			reservations[i].PaymentStatus = "Refunded"
			utils.WriteJSON("data/reservations.json", reservations)
			return nil
		}
	}
	return fmt.Errorf("Error: Reservation not found.")
}

func UpdateReservation(id int, roomID int, totalPayment int) error {
	for i, reservation := range reservations {
		if reservation.ID == id {
			reservations[i].RoomID = roomID
			reservations[i].TotalPayment = totalPayment
			utils.WriteJSON("data/reservations.json", reservations)
			return nil
		}
	}
	return fmt.Errorf("Error: Reservation not found.")
}

func MakeReservation(customerName, phoneNumber string, roomID int, reservationDate time.Time, totalPayment int) error {
    var reservation models.Reservation
    reservation.CustomerName = customerName
    reservation.CustomerPhone = phoneNumber
    reservation.RoomID = roomID
    reservation.TotalPayment = totalPayment
    reservation.PaymentStatus = "Pending"
    reservation.ReservationStatus = "Pending Approval"
    reservation.ReservationDate = reservationDate.Format("2006-01-02") 
    reservation.CreatedAt = time.Now().Format(time.RFC3339)

    reservation.ID = len(reservations) + 1
    reservations = append(reservations, reservation)

    err := utils.WriteJSON("data/reservations.json", reservations)
    if err != nil {
        return fmt.Errorf("Failed to save reservation: %v", err)
    }

    return nil
}

func CancelReservation(id int) error {
	for i, reservation := range reservations {
		if reservation.ID == id && reservation.ReservationStatus == "Pending Approval" {
			reservations = append(reservations[:i], reservations[i+1:]...)
			utils.WriteJSON("data/reservations.json", reservations)
			return nil
		}
	}
	return fmt.Errorf("Error canceling reservation: reservation not found or already confirmed")
}

func ViewMyReservations(customerName string) []models.Reservation {
	var myReservations []models.Reservation

	for _, r := range reservations {
		if r.CustomerName == customerName {
			myReservations = append(myReservations, r)
		}
	}
	return myReservations
}
