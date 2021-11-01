package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getOccupancyLevel(occupancyRate int) string {
	if occupancyRate < 30 {
		return "Low"
	} else if occupancyRate < 60 {
		return "Medium"
	} else {
		return "High"
	}
}

func getOccupancyLevelSwitch(occupancyRate int) string {

	switch {
	case occupancyRate < 30:
		return "Low"
	case occupancyRate < 60:
		return "Medium"
	default:
		return "High"
	}
}

func getOccupancyRate(roomsOccupied, nRooms int) int {
	return int(roomsOccupied * 100 / nRooms)
}

func printRoomDetails(roomNumber, nPeople, nNights int) {
	fmt.Printf("- %d : %d people / %d nights \n", roomNumber, nPeople, nNights)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const nRooms int = 134

	var roomsOccupied int = rand.Intn(nRooms)
	var roomAvailable int = nRooms - roomsOccupied
	occupancyRate := getOccupancyRate(roomsOccupied, nRooms)

	occupancyLevel := getOccupancyLevelSwitch(occupancyRate)

	fmt.Println("Gopher Paris Inn")
	fmt.Printf("Number of rooms %d \n", nRooms)
	fmt.Printf("Rooms available %d \n", roomAvailable)
	fmt.Printf("Occupancy Level %s \n", occupancyLevel)
	fmt.Println("Occupancy rate: ", occupancyRate)
	roomNumber := 110

	if roomAvailable > 0 {
		fmt.Println("Rooms:")
		for i := 1; i <= roomAvailable; i++ {
			var nPeople, nNights int = rand.Intn(5) + 1, rand.Intn(10) + 1
			printRoomDetails(roomNumber, nPeople, nNights)
			roomNumber++
		}
	} else {
		fmt.Println("No rooms Available for tonight ")
	}

}
