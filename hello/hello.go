package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message, err := greetings.Hello("Charlie")

	// If error encountered, log the error and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// Get a greeting for lots of people
	names := []string{"charlie", "erika"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	// If no error encountered, print the message
	fmt.Println(message)
	fmt.Println(messages)
}
