package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	readFile, err := os.Open("./input.txt")
	check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	max := 0
	current := 0

	for fileScanner.Scan() {
		value, err := strconv.Atoi(fileScanner.Text())

		if err == nil { // value is an int
			current += value
		} else { // value is a newline, so we stop counting for the current elf and start at 0 for next elf
			if current > max {
				max = current
			}

			current = 0
		}
	}

	fmt.Println(max)
}
