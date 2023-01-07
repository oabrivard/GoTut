package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const WINDOW_SIZE = 14

func shiftWindow(window []byte, b byte, counters *[26]int) {
	counters[window[0]]--
	for i := 0; i < WINDOW_SIZE-1; i++ {
		window[i] = window[i+1]

	}
	window[WINDOW_SIZE-1] = b
	counters[window[WINDOW_SIZE-1]]++
}

func main() {
	fileName := os.Args[1]

	bytes, err := os.ReadFile(fileName)
	check(err)

	window := make([]byte, WINDOW_SIZE)
	copy(window, bytes)

	for i := 0; i < WINDOW_SIZE; i++ {
		window[i] = window[i] - 'a'
	}

	fmt.Println(window)
	var counters [26]int

	for i := 0; i < WINDOW_SIZE; i++ {
		counters[window[i]]++
	}

	isMarker := true

	for _, counter := range counters {
		if counter > 1 {
			isMarker = false
		}
	}

	if isMarker {
		fmt.Println(WINDOW_SIZE)
		return
	}

	for i := WINDOW_SIZE; i < len(bytes); i++ {
		if bytes[i] >= 'a' && bytes[i] <= 'z' {
			shiftWindow(window, bytes[i]-'a', &counters)
		}

		isMarker = true

		for _, counter := range counters {
			if counter > 1 {
				isMarker = false
			}
		}

		if isMarker {
			fmt.Println(i + 1)
			return
		}
	}

	fmt.Println(counters)
}
