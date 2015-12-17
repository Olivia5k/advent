package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// Coord are coordinates on the grid
type Coord struct {
	x int
	y int
}

func main() {
	var presents map[Coord]int
	presents = make(map[Coord]int)

	santaX := 0
	santaY := 0
	roboX := 0
	roboY := 0

	presents[Coord{0, 0}] += 2
	data, err := ioutil.ReadFile("input/3")
	if err != nil {
		log.Fatal(err)
	}
	for i, c := range data {
		if i%2 == 0 {
			switch c {
			case 94: // Up
				santaY++
			case 118: // Down
				santaY--
			case 60: // Left
				santaX--
			case 62: // Right
				santaX++
			}
			presents[Coord{santaX, santaY}]++

		} else {
			switch c {
			case 94: // Up
				roboY++
			case 118: // Down
				roboY--
			case 60: // Left
				roboX--
			case 62: // Right
				roboX++
			}
			presents[Coord{roboX, roboY}]++
		}

	}

	fmt.Println(len(presents))
}
