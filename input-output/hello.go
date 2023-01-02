package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Print("What is your name? ")
	var name string
	if _, err := fmt.Scanln(&name); err != nil {
		log.Fatal(err)
	}

	result := fmt.Sprintf("Hello, %s, nice to meet you!", name)
	fmt.Println(result)
}
