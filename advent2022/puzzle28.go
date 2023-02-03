package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const CAVE_WIDTH = 1002

type coord struct {
	x int
	y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func lineToCoords(line string) []coord {
	result := []coord{}

	xys := strings.Split(line, " -> ")
	for i := range xys {
		xy := strings.Split(xys[i], ",")
		x, err := strconv.Atoi(xy[0])
		check(err)
		y, err := strconv.Atoi(xy[1])
		check(err)
		result = append(result, coord{x, y})
	}

	return result
}

func createLine(cave [][]rune, from *coord, to *coord) {
	//fmt.Println("Create line from", *from, "to", *to)

	if from.x == to.x {
		if from.y < to.y {
			for i := from.y + 1; i < to.y; i++ {
				cave[from.x][i] = '#'
			}
		} else {
			for i := to.y + 1; i < from.y; i++ {
				cave[from.x][i] = '#'
			}
		}
	} else if from.y == to.y {
		if from.x < to.x {
			for i := from.x + 1; i < to.x; i++ {
				cave[i][from.y] = '#'
			}
		} else {
			for i := to.x + 1; i < from.x; i++ {
				cave[i][from.y] = '#'
			}
		}
	} else {
		panic("Can't draw line in diagonal")
	}
}

// Parse input file to generate the matrix representing the cave structure
func parseMatrix(fileName string, maxDepth int) [][]rune {
	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

	cave := make([][]rune, CAVE_WIDTH)
	for i := range cave {
		cave[i] = make([]rune, maxDepth)

		for j := range cave[i] {
			cave[i][j] = ' '
		}
	}

	for y := 1; y < maxDepth; y++ {
		cave[0][y] = '|'
		cave[CAVE_WIDTH-1][y] = '|'
	}

	for x := 1; x < CAVE_WIDTH; x++ {
		cave[x][maxDepth-1] = '-'
	}

	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		coords := lineToCoords(line)

		var previous *coord
		for i, c := range coords {
			cave[c.x][c.y] = '#'

			if previous != nil {
				createLine(cave, previous, &c)
			}

			previous = &coords[i]
		}
	}

	return cave
}

// Find minimum amd maximum depth and width of cave
func findMinMaxDepthWidth(fileName string) (int, int, int) {
	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	maxDepth := 0
	minWidth := math.MaxInt
	maxWidth := 0
	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		coords := lineToCoords(line)

		for _, coord := range coords {
			if coord.x < minWidth {
				minWidth = coord.x
			}
			if coord.x > maxWidth {
				maxWidth = coord.x
			}
			if coord.y > maxDepth {
				maxDepth = coord.y
			}
		}
	}

	return minWidth, maxWidth, maxDepth
}

func printCave(cave [][]rune, minWidth int, maxWidth int, maxDepth int) {
	for y := 0; y < maxDepth; y++ {
		for x := minWidth - 10; x < maxWidth+10; x++ {
			fmt.Print(string(cave[x][y]))
		}
		fmt.Print("\n")
	}
}

func simFlow(cave [][]rune, maxDepth int) int {
	start := coord{500, -1}
	//cave[start.x][start.y] = '+'
	count := 0

	for {
		sandUnit := start
		if cave[sandUnit.x][sandUnit.y+1] == 'o' {
			return count
		}

		for {
			if cave[sandUnit.x][sandUnit.y+1] == ' ' {
				sandUnit.y++
				continue
			} else if cave[sandUnit.x-1][sandUnit.y+1] == ' ' {
				sandUnit.x--
				sandUnit.y++
				continue
			} else if cave[sandUnit.x+1][sandUnit.y+1] == ' ' {
				sandUnit.x++
				sandUnit.y++
				continue
			} else {
				// sand is stuck
				if sandUnit.y > maxDepth {
					panic("Unexpected state")
				} else {
					cave[sandUnit.x][sandUnit.y] = 'o'
					count++
					break
				}
			}
		}
	}
}

func main() {
	fileName := os.Args[1]

	minWidth, maxWidth, maxDepth := findMinMaxDepthWidth(fileName)
	maxDepth += 3
	fmt.Println(maxDepth)
	cave := parseMatrix(fileName, maxDepth)
	printCave(cave, minWidth, maxWidth, maxDepth)
	result := simFlow(cave, maxDepth)
	printCave(cave, minWidth, maxWidth, maxDepth)
	fmt.Println(result)
}
