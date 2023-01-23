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
	cycles := 0
	x := 1
	nextCheckedCycle := 20
	signalStrengh := 0

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		fields := strings.Fields(line)

		if fields[0] == "noop" {
			cycles++
			if cycles == nextCheckedCycle {
				signalStrengh += cycles * x
				nextCheckedCycle += 40
			}
		} else if fields[0] == "addx" {
			cycles++
			if cycles == nextCheckedCycle {
				signalStrengh += cycles * x
				nextCheckedCycle += 40
			}
			cycles++
			if cycles == nextCheckedCycle {
				signalStrengh += cycles * x
				nextCheckedCycle += 40
			}
			v, _ := strconv.Atoi(fields[1])
			x += v

		} else {
			panic(fmt.Sprintf("invalid instruction %s", fields[0]))
		}
	}

	fmt.Println(x)
	fmt.Println(signalStrengh)
}
