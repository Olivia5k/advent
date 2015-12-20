package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

var input = "ckczppom"

func main() {
	for x := 1; true; x++ {
		if x%10000 == 0 {
			fmt.Printf(".")
		}
		out := []byte(input + strconv.Itoa(x))
		hash := md5.Sum(out)
		s := fmt.Sprintf("%x", hash)
		if s[0:6] == "000000" {
			fmt.Println()
			fmt.Println(s)
			fmt.Println(x)
			break
		}
	}
}
