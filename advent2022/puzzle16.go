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

	fmt.Println(trees)

	height := len(trees)
	width := len(trees[0])

	max := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {

			// look rigth
			countR := 0
			for col := j + 1; col < width; col++ {
				countR++
				if trees[i][col] >= trees[i][j] {
					break
				}
			}

			// look left
			countL := 0
			for col := j - 1; col >= 0; col-- {
				countL++
				if trees[i][col] >= trees[i][j] {
					break
				}
			}

			// look down
			countD := 0
			for row := i + 1; row < height; row++ {
				countD++
				if trees[row][j] >= trees[i][j] {
					break
				}
			}

			// look up
			countU := 0
			for row := i - 1; row >= 0; row-- {
				countU++
				if trees[row][j] >= trees[i][j] {
					break
				}
			}

			count := countU * countR * countD * countL
			if count > max {
				max = count
				fmt.Println(trees[i][j])
				fmt.Println(i, j)
				fmt.Println("Count up", countU)
				fmt.Println("Count right", countR)
				fmt.Println("Count down", countD)
				fmt.Println("Count left", countL)
				fmt.Println("Count", count)
			}
		}
	}

	fmt.Println(max)

}
