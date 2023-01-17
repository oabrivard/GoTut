package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

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

	head := coord{0, 0}
	tail := coord{0, 0}

	positions := map[coord]bool{}
	positions[tail] = true

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		fields := strings.Fields(line)
		shift, _ := strconv.Atoi(fields[1])

		fmt.Println(line)
		for i := 0; i < shift; i++ {
			switch fields[0] {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y++
			case "D":
				head.y--
			default:
				panic("Invalid direction")
			}

			fmt.Println("H", head, "/ tail", tail)

			if head == tail {
				continue
			}

			around := [8]coord{
				{head.x, head.y - 1},
				{head.x, head.y + 1},
				{head.x - 1, head.y},
				{head.x + 1, head.y},
				{head.x - 1, head.y - 1},
				{head.x - 1, head.y + 1},
				{head.x + 1, head.y + 1},
				{head.x + 1, head.y - 1},
			}

			isClose := false
			for _, c := range around {
				if tail == c {
					isClose = true
				}
			}

			if isClose {
				continue
			}

			// Need to move tail

			if head.x == tail.x {
				if head.y > tail.y {
					tail.y = head.y - 1
				} else {
					tail.y = head.y + 1
				}

				positions[tail] = true
				fmt.Println("H", head, "/ tail", tail)
				continue
			}

			if head.y == tail.y {
				if head.x > tail.x {
					tail.x = head.x - 1
				} else {
					tail.x = head.x + 1
				}

				positions[tail] = true
				fmt.Println("H", head, "/ tail", tail)
				continue
			}

			if head.y == tail.y+2 {
				tail.y = head.y - 1
				tail.x = head.x
				positions[tail] = true
				fmt.Println("H", head, "/ tail", tail)
				continue
			}

			if head.x == tail.x+2 {
				tail.x = head.x - 1
				tail.y = head.y
				positions[tail] = true
				fmt.Println("H", head, "/ tail", tail)
				continue
			}

			if head.y == tail.y-2 {
				tail.y = head.y + 1
				tail.x = head.x
				positions[tail] = true
				fmt.Println("H", head, "/ tail", tail)
				continue
			}

			if head.x == tail.x-2 {
				tail.x = head.x + 1
				tail.y = head.y
				positions[tail] = true
				fmt.Println("H", head, "/ tail", tail)
				continue
			}

		}

	}

	fmt.Println(positions)
	fmt.Println(len(positions))

}
