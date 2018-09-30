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

func recursive(N int) {
	for i := 1; i <= N; i++ {
		recursive(N - 1)
	}
}

func main() {
	// floorCaterpillar := []int{1, 2, 3, 4}
	tableCaterpillar := []int{0, 0, 0, 0}
	var oneSegPoss int
	// Search number of ways to have M=1
	for i := 1; i <= 4; i++ { //get first card
		if i == 1 {
			tableCaterpillar[i-1] = i
			tableCaterpillar[i] = i + 1
			tableCaterpillar[i+1] = i + 2
			tableCaterpillar[i+2] = i + 3
			oneSegPoss += 1
			fmt.Println(tableCaterpillar, oneSegPoss)
		}
		if i == 2 {
			tableCaterpillar[0] = i
			for i := 1; i <= 4; i++ { //get second card
				if i == 2 || i == 4 { //second card cannot be the same as the first or 4 since that would create another segment
					continue
				} else {
					tableCaterpillar[1] = i
					secondCard := i
					for i := 1; i <= 4; i++ { //get third card
						if secondCard == 1 {
							if i == 2 || i == secondCard || i == 4 { //third card cannot be the same as the first of second
								continue
							} else {
								tableCaterpillar[2] = i
								thirdCard := i
								for i := 1; i <= 4; i++ { //get forth card
									if i == 2 || i == secondCard || i == thirdCard { //forth card cannot be the same as the first, second of third
										continue
									} else {
										tableCaterpillar[3] = i
										oneSegPoss += 1
										fmt.Println(tableCaterpillar, oneSegPoss)
									}
								}
							}
						} else {
							if i == 2 || i == secondCard { //third card cannot be the same as the first of second
								continue
							} else {
								tableCaterpillar[2] = i
								thirdCard := i
								for i := 1; i <= 4; i++ { //get forth card
									if i == 2 || i == secondCard || i == thirdCard { //forth card cannot be the same as the first, second of third
										continue
									} else {
										tableCaterpillar[3] = i
										oneSegPoss += 1
										fmt.Println(tableCaterpillar, oneSegPoss)
									}
								}
							}
						}
					}
				}
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
