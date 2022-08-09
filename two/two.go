package main

import (
	"fmt"

	"example.com/one"

	"log"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// // Get a greeting message and print it.
	// message := one.Hello("Gladys")
	// fmt.Println(message)

	// Request a greeting message.
	message, err := one.Hello("")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
