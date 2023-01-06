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

func shiftWindow(window []byte, b byte, counters *[26]int) {
	counters[window[0]]--
	window[0] = window[1]
	window[1] = window[2]
	window[2] = window[3]
	window[3] = b
	counters[window[3]]++
}

func main() {
	fileName := os.Args[1]

	bytes, err := os.ReadFile(fileName)
	check(err)

	window := make([]byte, 4)
	copy(window, bytes)

	window[0] = window[0] - 'a'
	window[1] = window[1] - 'a'
	window[2] = window[2] - 'a'
	window[3] = window[3] - 'a'

	var counters [26]int

	counters[window[0]]++
	counters[window[1]]++
	counters[window[2]]++
	counters[window[3]]++

	isMarker := true

	for _, counter := range counters {
		if counter > 1 {
			isMarker = false
		}
	}

	if isMarker {
		fmt.Println(4)
		return
	}

	for i := 4; i < len(bytes); i++ {
		shiftWindow(window, bytes[i]-'a', &counters)

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

	fmt.Println(window)
}
