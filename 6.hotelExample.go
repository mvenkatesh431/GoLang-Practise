package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const nRooms int = 134

	var roomsOccupied int = rand.Intn(nRooms)
	var roomAvailable int = nRooms - roomsOccupied
	occupancyRate := int(roomsOccupied * 100 / nRooms)

	var occupancyLevel string
	if occupancyRate < 30 {
		occupancyLevel = "Low"
	} else if occupancyRate < 60 {
		occupancyLevel = "Medium"
	} else {
		occupancyLevel = "High"
	}
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
			fmt.Printf("- %d : %d people / %d nights \n", roomNumber, nPeople, nNights)
			roomNumber++
		}
	} else {
		fmt.Println("No rooms Available for tonight ")
	}

}
