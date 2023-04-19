package main

import (
	"fmt"
	"github.com/yohann69/Go-Tutorials/greetings"
)

func main() {
	// Get a greeting message and print it
	message := greetings.Hello("Yohann")
	fmt.Println(message)
}
