package main

import (
	"bufio"
	"fmt"
	"os"
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

	items := []rune{}

	for fileScanner.Scan() {
		line1 := fileScanner.Text()
		fileScanner.Scan()
		line2 := fileScanner.Text()
		fileScanner.Scan()
		line3 := fileScanner.Text()

		for _, item := range line1 {
			if strings.ContainsRune(line2, item) && strings.ContainsRune(line3, item) {
				items = append(items, item)
				break // case where the shared item is present more than once
			}
		}
	}

	sum := 0
	for _, r := range items {
		if r >= 'A' && r <= 'Z' {
			sum += int(r - 'A' + 27)
		} else {
			sum += int(r - 'a' + 1)
		}
	}

	fmt.Println(sum)
	fmt.Println(string(items))
}
