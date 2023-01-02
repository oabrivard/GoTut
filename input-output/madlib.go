package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a noun: ")
	if !s.Scan() {
		log.Fatal(s.Err())
	}
	noun := s.Text()

	fmt.Print("Enter a verb: ")
	if !s.Scan() {
		log.Fatal(s.Err())
	}
	verb := s.Text()

	fmt.Print("Enter an adjective: ")
	if !s.Scan() {
		log.Fatal(s.Err())
	}
	adj := s.Text()

	fmt.Print("Enter an adverb: ")
	if !s.Scan() {
		log.Fatal(s.Err())
	}
	adv := s.Text()

	fmt.Printf("Do you %s your %s %s %s? That's hilarious!\n", verb, adj, noun, adv)
}
