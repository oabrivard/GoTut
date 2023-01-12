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

	trees := [][]int{}

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		treeLine := make([]int, len(line))

		for i, c := range line {
			treeLine[i] = int(c - '0')
		}

		trees = append(trees, treeLine)
	}

	//fmt.Println(trees)

	height := len(trees)
	width := len(trees[0])

	visibles := make([][]bool, height)
	for i := 0; i < height; i++ {
		visibles[i] = make([]bool, width)
		visibles[i][0] = true
		visibles[i][width-1] = true
	}

	for i := 0; i < width; i++ {
		visibles[0][i] = true
		visibles[height-1][i] = true
	}

	// scan each row from left to right
	for row := 1; row < height-1; row++ {
		max := trees[row][0]

		for col := 1; col < width-1; col++ {
			if trees[row][col] > max {
				visibles[row][col] = true
				max = trees[row][col]
			}
		}
	}

	// scan each row from right to left
	for row := 1; row < height-1; row++ {
		max := trees[row][width-1]

		for col := width - 2; col > 0; col-- {
			if trees[row][col] > max {
				visibles[row][col] = true
				max = trees[row][col]
			}
		}
	}

	// scan each columns from top to bottom
	for col := 1; col < width-1; col++ {
		max := trees[0][col]

		for row := 1; row < height-1; row++ {
			if trees[row][col] > max {
				visibles[row][col] = true
				max = trees[row][col]
			}
		}
	}

	// scan each columns from bottom to top
	for col := 1; col < width-1; col++ {
		max := trees[height-1][col]

		for row := height - 2; row > 0; row-- {
			if trees[row][col] > max {
				visibles[row][col] = true
				max = trees[row][col]
			}
		}
	}

	//fmt.Println(visibles)

	count := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if visibles[i][j] {
				count++
			}
		}
	}

	fmt.Println(count)

}
