package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"strings"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file and line number.
	log.SetPrefix(("greetings: "))
	log.SetFlags(50)
	// log.SetFlags(0)

	// // A slice of names.
	// names := []string{"Gladys", "Samantha", "Darrin"}

	// // Request a greeting message.
	// messages, err := greetings.Hellos(names)
	// // If an error was returned, print it to the console and
	// // exit the program.
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // If no error was returned, print the returned message
	// // to the console.
	// fmt.Println(messages)

	// Prompt the user to enter their name.
	fmt.Print("Enter your name: ")

	// Read user input from standard input (console).
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	// Trim any leading/trailing whitespace and newline characters
	// from the input
	name = strings.TrimSpace(name)

	// Greet the user
	message, err := greetings.Hello(name)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(message)
	}
}