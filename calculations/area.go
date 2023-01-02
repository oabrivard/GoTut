package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const SquareFeetToSquareMeterFactor = float64(0.09290304)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	fmt.Print("What is the length of the room in feet? ")
	if !s.Scan() {
		log.Fatal(s.Err())
	}

	length, err := strconv.ParseFloat(s.Text(), 64)
	checkErr(err)

	fmt.Print("What is the width of the room in feet? ")
	if !s.Scan() {
		log.Fatal(s.Err())
	}

	width, err := strconv.ParseFloat(s.Text(), 64)
	checkErr(err)

	sqFeet := length * width
	sqMeter := sqFeet * SquareFeetToSquareMeterFactor

	fmt.Printf(
		"You entered dimensions of %f feet by %f feet.\n"+
			"The area is %.3f square feet\n"+
			"equivalent to %.3f square meters\n",
		length, width, sqFeet, sqMeter)
}
