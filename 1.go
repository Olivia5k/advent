package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	x := 0
	first := 0
	data, err := ioutil.ReadFile("input/1")
	if err != nil {
		log.Fatal(err)
	}
	for i, c := range data {
		if c == 40 {
			x++
		} else {
			x--
			if first == 0 && x == -1 {
				first = i + 1
			}
		}
	}
	fmt.Println(x)     // First problem
	fmt.Println(first) // Second problem
}
