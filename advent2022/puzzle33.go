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

type rock struct {
	x      int
	y      int
	width  int
	height int
	shape  [][]int
}

func (r *rock) draw(tower *[HEIGHT][WIDTH]int) {
	for j := range r.shape {
		for i, val := range r.shape[j] {
			if val == 1 {
				tower[r.y+j][r.x+i] = 1
			}
		}
	}
}

func (r *rock) erase(tower *[HEIGHT][WIDTH]int) {
	for j := range r.shape {
		for i, val := range r.shape[j] {
			if val == 1 {
				tower[r.y+j][r.x+i] = 0
			}
		}
	}
}

func (r *rock) hasOverlapped(newX int, newY int, tower *[HEIGHT][WIDTH]int) bool {
	oldX := r.x
	oldY := r.y
	defer func() { r.x = oldX; r.y = oldY }()

	r.x = newX
	r.y = newY

	for j := range r.shape {
		for i, val := range r.shape[j] {
			if val == 1 {
				if tower[r.y+j][r.x+i] == 1 && val == 1 {
					return true
				}
			}
		}
	}

	return false
}

func (r *rock) appear(lastHeight int, tower *[HEIGHT][WIDTH]int) {
	r.x = 2
	r.y = lastHeight + 4

	r.draw(tower)
}

func (r *rock) push(direction rune, tower *[HEIGHT][WIDTH]int) {
	deltaX := 0

	if direction == '>' {
		if r.x+r.width < WIDTH {
			deltaX = 1
		} else {
			return
		}
	}

	if direction == '<' {
		if r.x > 0 {
			deltaX = -1
		} else {
			return
		}
	}

	r.erase(tower)
	if !r.hasOverlapped(r.x+deltaX, r.y, tower) {
		r.x += deltaX
	}
	r.draw(tower)
}

func (r *rock) fall(tower *[HEIGHT][WIDTH]int) (bool, int) {

	if r.y == 0 || r.hasOverlapped(r.x, r.y-1, tower) {
		return true, r.y
	}

	r.erase(tower)
	r.y--
	r.draw(tower)

	return false, r.y
}

func createRock(rockCount int) *rock {
	switch rockCount % 5 {
	case 0, 1, 2, 3, 4:
		return &rock{
			width:  4,
			height: 1,
			shape:  [][]int{{1, 1, 1, 1}},
		}
	default:
		panic("Impossible branch")
	}
}

func printTower(lastHeight int, tower *[HEIGHT][WIDTH]int) {
	for j := lastHeight; j >= 0; j-- {
		for _, val := range tower[j] {
			if val == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("-------")
	fmt.Print("\n")
}

const MAX_ROCKS = 4 //2022
const HEIGHT = (MAX_ROCKS + 1) * 4
const WIDTH = 7

func main() {
	fileName := os.Args[1]

	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Scan()
	streamJets := []rune(strings.TrimSpace(fileScanner.Text()))

	var tower [HEIGHT][WIDTH]int
	lastHeight := -1
	isBlocked := true
	var currentRock *rock = nil
	moves := 0

	for count := 0; count < MAX_ROCKS; {

		if isBlocked {
			isBlocked = false

			currentRock = createRock(count)
			currentRock.appear(lastHeight, &tower)
			count++
		}

		if moves%2 == 0 {
			d := streamJets[count%len(streamJets)]
			fmt.Println("push", string(d))
			currentRock.push(d, &tower)
		} else {
			fmt.Println("fall")
			isBlocked, lastHeight = currentRock.fall(&tower)
		}

		printTower(lastHeight+4, &tower)
		moves++
	}

	fmt.Println("Finished!")
}
