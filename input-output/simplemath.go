package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	fmt.Print("What is the first number? ")
	if !s.Scan() {
		log.Fatal(s.Err())
	}
	firstInput := s.Text()
	first, err := strconv.Atoi(firstInput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("What is the second number? ")
	if !s.Scan() {
		log.Fatal(s.Err())
	}
	secondInput := s.Text()
	second, err := strconv.Atoi(secondInput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(
		"%d + %d = %d\n"+
			"%d - %d = %d\n"+
			"%d * %d = %d\n"+
			"%d / %d = %.2f",
		first, second, first+second,
		first, second, first-second,
		first, second, first*second,
		first, second, float64(first)/float64(second))

}
