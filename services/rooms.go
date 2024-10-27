package services

import (
	"errors"
	"main/models"
	"main/utils"
)

var Rooms []models.Room

func init() {
	err := utils.ReadJSON("data/rooms.json", &Rooms)
	if err != nil {
		panic("Could not load room data.")
	}
}

func AddRoom(room models.Room) error {
	if room.Name == "" || room.Price <= 0 {
		return errors.New("room name and price are required")
	}
	room.ID = len(Rooms) + 1
	Rooms = append(Rooms, room)
	return utils.WriteJSON("data/rooms.json", Rooms)
}

func DeleteRoom(id int) error {
	for i, r := range Rooms {
		if r.ID == id {
			Rooms = append(Rooms[:i], Rooms[i+1:]...)
			return utils.WriteJSON("data/rooms.json", Rooms)
		}
	}
	return errors.New("Rooms not found")
}
