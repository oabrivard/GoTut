package main

import (
	"bufio"
	"fmt"
	"math"
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

func createLine(cave [][]rune, from *coord, to *coord, minWidth int) {
	//fmt.Println("Create line from", *from, "to", *to)

	if from.x == to.x {
		if from.y < to.y {
			for i := from.y + 1; i < to.y; i++ {
				cave[from.x-minWidth+1][i] = '#'
			}
		} else {
			for i := to.y + 1; i < from.y; i++ {
				cave[from.x-minWidth+1][i] = '#'
			}
		}
	} else if from.y == to.y {
		if from.x < to.x {
			for i := from.x + 1; i < to.x; i++ {
				cave[i-minWidth+1][from.y] = '#'
			}
		} else {
			for i := to.x + 1; i < from.x; i++ {
				cave[i-minWidth+1][from.y] = '#'
			}
		}
	} else {
		panic("Can't draw line in diagonal")
	}
}

// Parse input file to generate the matrix representing the cave structure
func parseMatrix(fileName string, minWidth int, maxWidth int, maxDepth int) [][]rune {
	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

	width := maxWidth - minWidth + 1

	cave := make([][]rune, width+2)
	for i := range cave {
		cave[i] = make([]rune, maxDepth+2)

		for j := range cave[i] {
			cave[i][j] = ' '
		}
	}

	for y := 1; y < maxDepth+2; y++ {
		cave[0][y] = '|'
		cave[width+1][y] = '|'
	}

	for x := 1; x <= width; x++ {
		cave[x][maxDepth+1] = '-'
	}

	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		coords := lineToCoords(line)

		var previous *coord
		for i, c := range coords {
			x := c.x - minWidth + 1
			y := c.y
			cave[x][y] = '#'

			if previous != nil {
				createLine(cave, previous, &c, minWidth)
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
	width := maxWidth - minWidth + 2

	for y := 0; y < maxDepth+2; y++ {
		for x := 0; x < width+1; x++ {
			fmt.Print(string(cave[x][y]))
		}
		fmt.Print("\n")
	}
}

func simFlow(cave [][]rune, minWidth int, maxWidth int, maxDepth int) int {
	start := coord{500 - minWidth + 2, 0}
	cave[start.x][start.y] = '+'
	count := 0

	for {
		sandUnit := start

		for {
			if cave[sandUnit.x][sandUnit.y+1] == ' ' {
				sandUnit.y++
				continue
			} else {
				if cave[sandUnit.x-1][sandUnit.y+1] == ' ' {
					sandUnit.x--
					sandUnit.y++
					continue
				} else if cave[sandUnit.x+1][sandUnit.y+1] == ' ' {
					sandUnit.x++
					sandUnit.y++
					continue
				} else if cave[sandUnit.x-1][sandUnit.y+1] == '|' {
					cave[sandUnit.x][sandUnit.y] = 'X'
					return count
				} else if cave[sandUnit.x+1][sandUnit.y+1] == '|' {
					cave[sandUnit.x][sandUnit.y] = 'X'
					return count
				} else {
					// sand is stuck
					if sandUnit.y > maxDepth {
						panic("Unexpected state")
					} else {
						if cave[sandUnit.x][sandUnit.y+1] == '#' {
							cave[sandUnit.x][sandUnit.y] = 'o'
							count++
							break
						} else {
							if cave[sandUnit.x-1][sandUnit.y] == '|' || cave[sandUnit.x+1][sandUnit.y] == '|' || cave[sandUnit.x][sandUnit.y+1] == '-' {
								cave[sandUnit.x][sandUnit.y] = 'X'
								return count
							}
							cave[sandUnit.x][sandUnit.y] = 'o'
							count++
							break
						}
					}
				}
			}

		}
	}

}

func main() {
	fileName := os.Args[1]

	minWidth, maxWidth, maxDepth := findMinMaxDepthWidth(fileName)
	fmt.Println(minWidth, maxWidth, maxDepth)
	cave := parseMatrix(fileName, minWidth, maxWidth, maxDepth)
	printCave(cave, minWidth, maxWidth, maxDepth)
	result := simFlow(cave, minWidth+1, maxWidth, maxDepth)
	printCave(cave, minWidth, maxWidth, maxDepth)
	fmt.Println(result)
}
