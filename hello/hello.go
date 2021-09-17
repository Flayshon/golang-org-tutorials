package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix( "greetings: ")
	log.SetFlags(0)

	greetings.Init()

	message, err := greetings.Hello("Flayshon")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}