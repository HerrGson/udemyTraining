package main

import (
	"fmt"
	"math/rand"
)

var nrOfJigSawPieces int = 10

func findNewLocation(caterpillar [10]int) int {
	newLocation := rand.Intn(10)
	i := 0
	for caterpillar[newLocation] != 0 {
		newLocation = rand.Intn(10)
		if i > 100 {
			break
		} else {
			i++
		}
	}
	return newLocation
}

func initCaterpillar() [10]int {
	var caterpillar [10]int
	for i := 1; i <= 10; i++ {
		caterpillar[findNewLocation(caterpillar)] = i
	}
	return caterpillar
}

func recursiveLoop(stage int) int {
	var nrOfIterationsRan int
	for i := 1; i <= stage; i++ {
		nrOfIterationsRan = stage * recursiveLoop(stage-1)
	}
	return nrOfIterationsRan
}

func cardStillOnFloor(tableCaterpillar []int, newCard int) bool {
	for _, value := range tableCaterpillar {
		if value == newCard {
			return false
		}
	}
	return true
}

func main() {
	// floorCaterpillar := initCaterpillar()
	// fmt.Println("The floor pieces: ", floorCaterpillar)
	N := 5
	var floorCaterpillar []int
	for i := 1; i <= N; i++ {
		floorCaterpillar = append(floorCaterpillar, i)
	}
	fmt.Println("floorCaterpillar: ", floorCaterpillar)
	var tableCaterpillar []int
	for i := 1; i <= N; i++ {
		tableCaterpillar = append(tableCaterpillar, 0)
	}
	fmt.Println("tableCaterpillar: ", tableCaterpillar)

	var sum int
	nrOfSegments := 1
	// var sumInSegments [][]int

	// Loop over all combinations
	for i := 1; i <= N; i++ {

		// initialize new table catepillar
		var tableCaterpillar []int
		for i := 1; i <= N; i++ {
			tableCaterpillar = append(tableCaterpillar, 0)
		}

		// Put newley pulled jigSaw on the table
		tableCaterpillar[i] = i

		// Get next jigSaw and put it on table
		for i := 1; i <= N; i++ {
			if cardStillOnFloor(tableCaterpillar, i) {
				tableCaterpillar = append(tableCaterpillar, i)
				if tableCaterpillar[i-1] == 0 && tableCaterpillar[i+1] == 0 {
					nrOfSegments += 1
				} else if tableCaterpillar[i-1] != 0 && tableCaterpillar[i+1] != 0 {
					nrOfSegments -= 1
				}
			}
		}
	}
	fmt.Println(sum)

	// Total nr of possibilities to put together puzzle with only ever 1 segment
	recursiveLoop(5)
	fmt.Println("recursive sum", sum)
	// test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(test[10])
}
