package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	fmt.Print("What is the quote? ")
	if !s.Scan() {
		log.Fatal(s.Err())
	}
	quote := s.Text()

	fmt.Print("Who said it? ")
	if !s.Scan() {
		log.Fatal(s.Err())
	}
	author := s.Text()

	fmt.Println(author, "says, \""+quote+"\"")
}
