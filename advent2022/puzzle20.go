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

func drawPixel(pixel int, x int) {
	if pixel == x-1 || pixel == x || pixel == x+1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if pixel == 39 {
		fmt.Println("")
	}
}

func main() {
	fileName := os.Args[1]

	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	x := 1
	pixel := 0

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		fields := strings.Fields(line)

		if fields[0] == "noop" {
			drawPixel(pixel, x)
			pixel = (pixel + 1) % 40
		} else if fields[0] == "addx" {
			drawPixel(pixel, x)
			pixel = (pixel + 1) % 40
			drawPixel(pixel, x)
			pixel = (pixel + 1) % 40
			v, _ := strconv.Atoi(fields[1])
			x += v
		} else {
			panic(fmt.Sprintf("invalid instruction %s", fields[0]))
		}
	}

	fmt.Println(x)
}
