package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Print("What is the input string? ")

	var s string
	if _, err := fmt.Scanln(&s); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s has %d characters.\n", s, len(s))
}
