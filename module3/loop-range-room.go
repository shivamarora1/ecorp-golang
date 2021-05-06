package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	type Room struct {
		CarpetArea            int //sqrtFt
		TotalPicturesCapacity int
		PicturesAvailable     []string
		WallColor             string
		RoomType              string
	}

	livingRoom := Room{CarpetArea: 230, TotalPicturesCapacity: 5,
		PicturesAvailable: []string{"Elephant", "Family"},
		WallColor:         "Red", RoomType: "Living Room"}
	frontBedRoom := Room{CarpetArea: 100, TotalPicturesCapacity: 8,
		PicturesAvailable: []string{"Emily"},
		WallColor:         "Red", RoomType: "Front Bed Room"}
	backBedRoom := Room{CarpetArea: 150, TotalPicturesCapacity: 6,
		PicturesAvailable: []string{}, WallColor: "Yellow", RoomType: "Back Bed Room"}
	kitchen := Room{CarpetArea: 50, TotalPicturesCapacity: 5,
		PicturesAvailable: []string{}, WallColor: "Red", RoomType: "Kitchen"}
	study := Room{CarpetArea: 50, TotalPicturesCapacity: 3,
		PicturesAvailable: []string{"Bill Gates", "Steve Jobs"},
		WallColor:         "Red", RoomType: "Study Room"}

	allImages := []string{"Sunder_Pichai_Quotes", "Mountains",
		"Snow_Man", "Celebration", "Family_Photo", "Marriage_Photo_Graph", "Jesus_Christ",
		"Thor", "Friends_Pictures", "Eagle", "Kingfisher"}

	house := []Room{livingRoom, frontBedRoom, backBedRoom, kitchen, study}

	for _, rm := range house {
		fmt.Printf("Before updation %+v\n", rm)
	}

	for i, rm := range house {
		availableImages := len(rm.PicturesAvailable)
		capacity := rm.TotalPicturesCapacity
		neededImgs := capacity - availableImages

		for j := 0; j < neededImgs; j++ {
			//Choose a random picture and insert into wall
			img := allImages[rand.Intn(len(allImages))]
			rm.PicturesAvailable = append(rm.PicturesAvailable, img)
		}

		house[i] = rm
	}

	//Printing results
	for _, rm := range house {
		fmt.Printf("%+v\n", rm)
	}

}
