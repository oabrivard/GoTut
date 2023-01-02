package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	fmt.Print("What is your current age? ")
	if !s.Scan() {
		log.Fatal(s.Err())
	}
	current, err := strconv.Atoi(s.Text())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("At what age would you like to retire? ")
	if !s.Scan() {
		log.Fatal(s.Err())
	}
	retirement, err := strconv.Atoi(s.Text())
	if err != nil {
		log.Fatal(err)
	}

	year := time.Now().Year()
	yearsLeft := retirement - current

	if yearsLeft <= 0 {
		fmt.Println("You can already retire.")
	} else {
		fmt.Printf("You have %d years left until you can retire.\nIt's %d, so you can retire in %d.\n",
			yearsLeft, year, year+yearsLeft)
	}
}
