package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	count := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		ranges := strings.Split(line, ",")
		range1 := strings.Split(ranges[0], "-")
		range2 := strings.Split(ranges[1], "-")
		lower1, _ := strconv.Atoi(range1[0])
		upper1, _ := strconv.Atoi(range1[1])
		lower2, _ := strconv.Atoi(range2[0])
		upper2, _ := strconv.Atoi(range2[1])

		// fmt.Println(lower1, upper1, lower2, upper2)

		if lower2 <= lower1 && upper1 <= upper2 || lower1 <= lower2 && upper2 <= upper1 {
			count++
		}
	}

	fmt.Println(count)
}
