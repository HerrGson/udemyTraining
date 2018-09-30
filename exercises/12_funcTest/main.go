package main

import "fmt"

func numberInSlice(number int, slice []int) bool {
	for _, value := range slice {
		if number == value {
			return true
		}
	}
	return false
}
func newSlice(number int, segments int, slice []int, emptyCaterpillar bool) bool {
	if emptyCaterpillar {
		return false
	}
	if number == 1 {
		for _, value := range slice {
			if number+1 == value {
				return false
			}
		}
	} else if number == len(slice) {
		for _, value := range slice {
			if number-1 == value {
				return false
			}
		}
	} else {
		neighboursUpstairs := false
		neighboursDownstairs := false
		for _, value := range slice {
			if number+1 == value {
				neighboursUpstairs = true
			} else if number-1 == value {
				neighboursDownstairs = true
			}
		}
		if !neighboursUpstairs && !neighboursDownstairs {
			return true
		} else if !neighboursUpstairs || !neighboursDownstairs {
			return false
		}
	}
	return true
}

func recursive(tableCaterpillar []int, currentLvl int, oneSegPoss int) {
	for i := 1; i <= 4; i++ {
		fmt.Println("checking index ", 3-(currentLvl-1), "with card", i)
		emptyCaterpillar := true
		for _, value := range tableCaterpillar {
			if value == 0 {
			} else {
				emptyCaterpillar = false
			}
		}
		fmt.Println("This is a new slice: ", newSlice(i, 1, tableCaterpillar, emptyCaterpillar))
		if numberInSlice(i, tableCaterpillar) || newSlice(i, 1, tableCaterpillar, emptyCaterpillar) { //second card cannot be the same as the first or 4 since that would create another segment
			continue
		} else {
			fmt.Println("update catepillar with ", i, "on index", 3-(currentLvl-1))
			tableCaterpillar[3-(currentLvl-1)] = i
			fmt.Println("catepillar updated to: ", tableCaterpillar)

			// Check if series is complete
			for k, value := range tableCaterpillar {
				if value == 0 {
					break
				} else if k+1 == len(tableCaterpillar) {
					oneSegPoss += 1
					fmt.Println("1 poss complete: ", tableCaterpillar, oneSegPoss)
					tableCaterpillar[3-(currentLvl-1)] = 0
					fmt.Println("after update: ", tableCaterpillar)
				} else {
					continue
				}
			}
			if currentLvl == 1 {
				continue
			}
			recursive(tableCaterpillar, currentLvl-1, oneSegPoss)
			fmt.Println("Went up one level to check index: ", 3-(currentLvl-1))
			fmt.Println("Here the caterpillar is: ", tableCaterpillar)

		}
	}
}

func main() {
	// floorCaterpillar := []int{1, 2, 3, 4}

	var oneSegPoss int
	tableCaterpillar := []int{0, 0, 0, 0}
	recursive(tableCaterpillar, 4, oneSegPoss)

	var seriesSum float64
	var possibilities float64
	var average float64
	seriesSum = 1*512 + 2*250912 + 3*1815264 + 4*1418112 + 5*144000
	possibilities = 512 + 250912 + 1815264 + 1418112 + 144000
	average = seriesSum / possibilities
	fmt.Println(average)
}
