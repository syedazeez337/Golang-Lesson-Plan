package main

import (
	"fmt"
	"math"
)

var (
	word string
	num  int
	pi   float64
	char byte
)

func main() {
	char = 'A'
	pi = math.Pi
	num = 23
	word = "Hello world"

	fmt.Printf("%c\n%.12f\n%d\n%s\n", char, pi, num, word)
}
