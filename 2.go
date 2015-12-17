package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// getArea returns the area of wrapping paper needed
func getArea(l, w, h int) int {
	x := 0
	lw := l * w * 2
	wh := w * h * 2
	hl := h * l * 2

	all := []int{lw, wh, hl}
	sort.Ints(all)
	low := all[0]

	x = lw + wh + hl + low/2

	return x
}

// getRibbon returns the ribbon length needed
func getRibbon(l, w, h int) int {
	y := 0

	all := []int{l, w, h}
	sort.Ints(all)

	y += all[0]*2 + all[1]*2
	y += l * w * h
	return y
}

func main() {
	file, err := os.Open("input/2")
	x := 0
	y := 0
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dim := strings.Split(line, "x")
		l, _ := strconv.Atoi(dim[0])
		w, _ := strconv.Atoi(dim[1])
		h, _ := strconv.Atoi(dim[2])

		x += getArea(l, w, h)
		y += getRibbon(l, w, h)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(x)
	fmt.Println(y)
}
