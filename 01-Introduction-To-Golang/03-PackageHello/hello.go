package hello

import (
	"fmt"
)

// Exported function
func SayHello() {
	fmt.Println("Hello world")
}

// Unexported function
func sayGoodBye() {
	fmt.Println("Goodbye")
}