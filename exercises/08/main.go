package main

import (
	"fmt"
	"math/rand"
)

func initCatepillar() [40]int {
	var catepillar [40]int
	for i := 1; i <= 40; i++ {
		catepillar[rand.Intn(39)] = i
	}
	return catepillar
}

func main() {
	floorCaterpillar := initCatepillar()
	var tableCaterpillar [40]int
	var nrOfSegments int
	var tableCaterpillar [40]int

	for i := 0; i <= 40; i++ {
		jigSawPiece := floorCaterpillar[i]
		tableCaterpillar[jigSawPiece] = jigSawPiece

		if jigSawPiece == 1 {
			if tableCaterpillar[2] == 0 {
				nrOfSegments += 1
			}
		} else if jigSawPiece == 40 {
			if tableCaterpillar[39] == 0 {
				nrOfSegments += 1
			}
		} else {
			if tableCaterpillar[jigSawPiece-1] == 0 && tableCaterpillar[jigSawPiece+1] == 0 {
				nrOfSegments += 1
			} else if tableCaterpillar[jigSawPiece-1] != 0 && tableCaterpillar[jigSawPiece+1] != 0 {
				nrOfSegments -= 1
			} else if tableCaterpillar[jigSawPiece-1] != 0 || tableCaterpillar[jigSawPiece+1] == 0 {
			}
		}

	}

	fmt.Println(floorCaterpillar)
	fmt.Println(tableCaterpillar)
}
