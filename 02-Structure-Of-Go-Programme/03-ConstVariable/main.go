package main

import (
	"fmt"
)

func main() {
	const num = 34

	var num2 = 23
	num2 = num2 + 1
	num3 := num + 1

	fmt.Println(num2, num3)
}