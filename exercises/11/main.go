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
func newSlice(number int, segments int, slice []int) bool {
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

func main() {
	// floorCaterpillar := []int{1, 2, 3, 4}

	var oneSegPoss int
	// Search number of ways to have M=1
	for i := 1; i <= 4; i++ { //get first card
		tableCaterpillar := []int{0, 0, 0, 0}
		tableCaterpillar[0] = i
		fmt.Println("First card is: ", i)
		for i := 1; i <= 4; i++ { //get second card
			if numberInSlice(i, tableCaterpillar) || newSlice(i, 1, tableCaterpillar) { //second card cannot be the same as the first or 4 since that would create another segment
				continue
			} else {
				tableCaterpillar[1] = i
				for i := 1; i <= 4; i++ { //get third card
					if numberInSlice(i, tableCaterpillar) || newSlice(i, 1, tableCaterpillar) { //third card cannot be the same as the first of second
						continue
					} else {
						tableCaterpillar[2] = i
						for i := 1; i <= 4; i++ { //get forth card
							if numberInSlice(i, tableCaterpillar) || newSlice(i, 1, tableCaterpillar) { //forth card cannot be the same as the first, second of third
								continue
							} else {
								tableCaterpillar[3] = i
								oneSegPoss += 1
								fmt.Println("1 poss complete: ", tableCaterpillar, oneSegPoss)
								tableCaterpillar[3] = 0
							}
						}
						tableCaterpillar[2] = 0
					}
				}
				tableCaterpillar[1] = 0
			}
		}
	}

	var seriesSum float64
	var possibilities float64
	var average float64
	seriesSum = 1*512 + 2*250912 + 3*1815264 + 4*1418112 + 5*144000
	possibilities = 512 + 250912 + 1815264 + 1418112 + 144000
	average = seriesSum / possibilities
	fmt.Println(average)
}
