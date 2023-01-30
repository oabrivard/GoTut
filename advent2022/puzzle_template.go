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

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		value, err := strconv.Atoi(line)
		check(err)
		fmt.Println(value)
	}

	fmt.Println("Finished!")
}
