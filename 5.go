package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// isNice determines if a string is nice or not
func isNice(s string) bool {
	return hasVowels(s) && hasDouble(s) && isWhitelisted(s)
}

func hasVowels(s string) bool {
	x := 0
	for _, c := range s {
		c := string(c)
		if strings.Contains("aeiou", c) {
			x++
			if x == 3 {
				return true
			}
		}
	}
	return false
}

func hasDouble(s string) bool {
	for _, c := range s {
		c := string(c)
		if strings.Contains(s, c+c) {
			return true
		}
	}
	return false
}

func isWhitelisted(s string) bool {
	for _, ok := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(s, ok) {
			return false
		}
	}
	return true
}

func isNewNice(s string) bool {
	return hasPair(s) && hasSurround(s)
}

func hasPair(s string) bool {
	for i, c := range s {
		if i+2 > len(s) {
			continue
		}
		c := string(c)
		n := string(s[i+1])
		// Make a new string that has the current pair removed
		t := s[0:i] + s[i+2:len(s)]
		if strings.Contains(t, c+n) {
			// fmt.Printf("pair: %s - %s\n", s, c+n)
			return true
		}
	}
	return false
}

func hasSurround(s string) bool {
	for i, c := range s {
		if i+3 > len(s) {
			continue
		}
		c := string(c)
		o := string(s[i+2])
		m := string(s[i+1])
		if c == o && c != m {
			// fmt.Printf("surround: %s - %s\n", s, c+m+o)
			return true
		}
	}
	return false
}

func main() {
	nice := 0
	newNice := 0

	file, err := os.Open("input/5")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isNice(line) {
			nice++
		}
		if isNewNice(line) {
			newNice++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(nice)
	fmt.Println(newNice)
}
