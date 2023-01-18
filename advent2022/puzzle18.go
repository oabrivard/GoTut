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

	rope := [10]coord{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}

	positions := map[coord]bool{}
	positions[rope[9]] = true

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		fields := strings.Fields(line)
		shift, _ := strconv.Atoi(fields[1])

		fmt.Println(line)

		for i := 0; i < shift; i++ {
			head := &rope[0]

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

			for j := 0; j < 9; j++ {
				head = &rope[j]
				tail := &rope[j+1]

				fmt.Println("Move H", j, "to", *head)

				if *head == *tail {
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
					if *tail == c {
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

					fmt.Println("aMove T", j+1, "to", *tail)

					if j+1 == 9 {
						positions[*tail] = true
					}
					continue
				}

				if head.y == tail.y {
					if head.x > tail.x {
						tail.x = head.x - 1
					} else {
						tail.x = head.x + 1
					}

					fmt.Println("bMove T", j+1, "to", *tail)

					if j+1 == 9 {
						positions[*tail] = true
					}
					continue
				}

				if head.x == tail.x+2 && head.y == tail.y+2 {
					tail.y = head.y - 1
					tail.x = head.x - 1

					fmt.Println("gMove T", j+1, "to", *tail)

					if j+1 == 9 {
						positions[*tail] = true
					}
					continue
				}

				if head.x == tail.x-2 && head.y == tail.y-2 {
					tail.y = head.y + 1
					tail.x = head.x + 1

					fmt.Println("hMove T", j+1, "to", *tail)

					if j+1 == 9 {
						positions[*tail] = true
					}
					continue
				}

				if head.x == tail.x+2 && head.y == tail.y-2 {
					tail.y = head.y + 1
					tail.x = head.x - 1

					fmt.Println("iMove T", j+1, "to", *tail)

					if j+1 == 9 {
						positions[*tail] = true
					}
					continue
				}

				if head.x == tail.x-2 && head.y == tail.y+2 {
					tail.y = head.y - 1
					tail.x = head.x + 1

					fmt.Println("jMove T", j+1, "to", *tail)

					if j+1 == 9 {
						positions[*tail] = true
					}
					continue
				}

				if head.y == tail.y+2 {
					tail.y = head.y - 1
					tail.x = head.x

					fmt.Println("cMove T", j+1, "to", *tail)

					if j+1 == 9 {
						positions[*tail] = true
					}
					continue
				}

				if head.x == tail.x+2 {
					tail.x = head.x - 1
					tail.y = head.y

					fmt.Println("dMove T", j+1, "to", *tail)

					if j+1 == 9 {
						positions[*tail] = true
					}
					continue
				}

				if head.y == tail.y-2 {
					tail.y = head.y + 1
					tail.x = head.x

					fmt.Println("eMove T", j+1, "to", *tail)

					if j+1 == 9 {
						positions[*tail] = true
					}
					continue
				}

				if head.x == tail.x-2 {
					tail.x = head.x + 1
					tail.y = head.y

					fmt.Println("fMove T", j+1, "to", *tail)

					if j+1 == 9 {
						positions[*tail] = true
					}
					continue
				}

			}
		}
	}

	fmt.Println(positions)
	fmt.Println(len(positions))

}
