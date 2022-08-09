package main

import (
	"fmt"

	"example.com/one"
)

func main() {
	// Get a greeting message and print it.
	message := one.Hello("Gladys")
	fmt.Println(message)
}
