package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileName := os.Args[1]

	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	scores := map[string]int{
		"A X": 3, // rock
		"A Y": 4,
		"A Z": 8,
		"B X": 1, // paper
		"B Y": 5,
		"B Z": 9,
		"C X": 2, // scisors
		"C Y": 6,
		"C Z": 7,
	}

	score := 0

	for fileScanner.Scan() {
		value := scores[fileScanner.Text()]
		score += value
	}

	fmt.Println(score)
}
