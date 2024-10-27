package views

import (
	"fmt"
	"main/models"
	"main/services"
	"main/utils"
	"strings"
)

func ManageRooms() {
	var action int
	fmt.Println("Room Management:")
	fmt.Println("1. Add room")
	fmt.Println("2. View rooms")
	fmt.Println("3. Edit room")
	fmt.Println("4. Delete room")
	fmt.Print("Select an action: ")
	fmt.Scanln(&action)

	switch action {
	case 1:
		AddRoomMenu()
	case 2:
		ViewRooms()
	case 3:
		UpdateRoomMenu()
	case 4:
		DeleteRoomMenu()
	default:
		fmt.Println("Invalid action")
	}
}

func AddRoomMenu() {
	var room models.Room
	name, err := utils.InputStr("Enter room Name: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	room.Name = name
	price, err := utils.InputInt("Enter room Price: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	room.Price = price
	isAvailable, err := utils.InputBool("Is Room Available (true/false): ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if isAvailable {
		room.IsAvailable = true
	} else {
		room.IsAvailable = false
	}

	room.IsAvailable = isAvailable

	var facilities []models.Facility
	for {
		var facility models.Facility
		name, err := utils.InputStr("Enter Facility Name (or type 'done' to finish): ")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		facility.Name = name
		if facility.Name == "done" {
			break
		}
		facility.ID = len(facilities) + 1
		facilities = append(facilities, facility)
	}

	room.Facilities = facilities

	if err := services.AddRoom(room); err != nil {
		fmt.Println("Error adding room:", err)
	} else {
		fmt.Println("room added successfully!")
	}
}

func ViewRooms() {
	fmt.Println("Rooms:")
	for _, r := range services.Rooms {
		availability := "Unavailable"
		if r.IsAvailable {
			availability = "Available"
		}
		fmt.Printf("ID: %d, Name: %s, Price: %d, Is_Available: %s\n", r.ID, r.Name, r.Price, availability)
		fmt.Println("Facilities:")
		for _, f := range r.Facilities {
			fmt.Printf("- %s\n", f.Name)
		}
	}
}

func ViewAvailableRooms() {
	availableRooms := services.Rooms

	fmt.Printf("\nAvailable Rooms\n")

	fmt.Println(strings.Repeat("-", 45))

	for _, r := range availableRooms {
		if r.IsAvailable {
			fmt.Printf("ID: %d, Name: %s, Price: %d\n", r.ID, r.Name, r.Price)
			for _, f := range r.Facilities {
				fmt.Printf("- %s\n", f.Name)
			}
			fmt.Println(strings.Repeat("-", 45))
		}
	}
}

func DeleteRoomMenu() {
	// var id int
	// fmt.Print("Enter room ID to delete: ")
	// fmt.Scanln(&id)

	id, err := utils.InputInt("Enter room ID to delete: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err := services.DeleteRoom(id); err != nil {
		fmt.Println("Error deleting room:", err)
	} else {
		fmt.Println("Room deleted successfully!")
	}
}

func UpdateRoomMenu() {
	var id int
	id, err := utils.InputInt("Enter room ID to update: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, r := range services.Rooms {
		if r.ID == id {

			name, err := utils.InputStr("Enter new room Name (leave blank to keep current): ")
			if err == nil && name != "" {
				r.Name = name
			}

			price, err := utils.InputInt("Enter new room Price (leave blank to keep current): ")
			if err == nil && price != 0 {
				r.Price = price
			}

			isAvailable, err := utils.InputBool("Is Room Available (leave blank to keep current, true/false): ")
			if err == nil {
				r.IsAvailable = isAvailable
			}

			fmt.Println("Current facilities:")
			for j, facility := range r.Facilities {
				fmt.Printf("%d. %s\n", j+1, facility.Name)
			}

			fmt.Println("Options:")
			fmt.Println("1. Edit a facility")
			fmt.Println("2. Add a new facility")
			fmt.Println("3. Remove a facility")
			fmt.Println("4. Keep current facilities")

			choice, err := utils.InputInt("Enter your choice: ")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			switch choice {
			case 1:
				facilityIndex, err := utils.InputInt("Enter the facility number to edit: ")
				if err == nil && facilityIndex > 0 && facilityIndex <= len(r.Facilities) {
					newName, err := utils.InputStr("Enter new name for the facility: ")
					if err == nil && newName != "" {
						r.Facilities[facilityIndex-1].Name = newName
					}
				} else {
					fmt.Println("Invalid facility number.")
				}
			case 2:
				newFacilityName, err := utils.InputStr("Enter new facility name: ")
				if err == nil && newFacilityName != "" {
					newFacility := models.Facility{
						ID:   len(r.Facilities) + 1,
						Name: newFacilityName,
					}
					r.Facilities = append(r.Facilities, newFacility)
				}
			case 3:
				facilityIndex, err := utils.InputInt("Enter the facility number to remove: ")
				if err == nil && facilityIndex > 0 && facilityIndex <= len(r.Facilities) {
					r.Facilities = append(r.Facilities[:facilityIndex-1], r.Facilities[facilityIndex:]...)
				} else {
					fmt.Println("Invalid facility number.")
				}
			case 4:

			default:
				fmt.Println("Invalid choice.")
			}

			services.Rooms[i] = r
			if err := utils.WriteJSON("data/rooms.json", services.Rooms); err != nil {
				fmt.Println("Error updating room:", err)
			} else {
				fmt.Println("Room updated successfully!")
			}
			return
		}
	}
	fmt.Println("Error updating room: room not found.")
}
