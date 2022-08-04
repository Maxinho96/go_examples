package main

import (
	"fmt"

	"github.com/Maxinho96/go_examples/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
