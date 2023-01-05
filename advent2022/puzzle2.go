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
	fileName := os.Args[1]

	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	max1, max2, max3 := 0, 0, 0
	current := 0

	for fileScanner.Scan() {
		value, err := strconv.Atoi(fileScanner.Text())

		if err == nil { // value is an int
			current += value
		} else { // value is a newline, so we stop counting for the current elf and start at 0 for next elf
			if current > max1 {
				max3 = max2
				max2 = max1
				max1 = current
			} else if current > max2 {
				max3 = max2
				max2 = current
			} else if current > max3 {
				max3 = current
			}

			current = 0
		}
	}

	fmt.Println(max1 + max2 + max3)
}
