package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message
	customMessage, err := greetings.Hello("Big D")
	// if an error was returned, print it to stdout and exit
	if err != nil {
		log.Fatal(err)
	}
	// ELSE print the returned message
	fmt.Println(customMessage)
}
