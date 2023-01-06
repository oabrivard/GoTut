package main

import (
	"bufio"
	"container/list"
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

	var stacks []*list.List

	// parse stacks
	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = line + " "
		count := len(line) / 4

		if line == " " {
			break
		}

		if stacks == nil {
			stacks = make([]*list.List, count)

			for idx := range stacks {
				stacks[idx] = list.New()
			}
		}

		fields := make([]string, count)
		for idx := range fields {
			field := strings.TrimSpace(line[idx*4 : idx*4+3])

			_, err := strconv.Atoi(field)

			if err != nil && field != "" {
				// here we parse "[X]" like strings
				field = strings.ReplaceAll(strings.ReplaceAll(field, "[", ""), "]", "")
				stacks[idx].PushBack(field)
			}

		}
	}

	// parse moves
	for fileScanner.Scan() {
		line := fileScanner.Text()

		fields := strings.Fields(line)
		if len(fields) < 6 {
			panic(line)
		}
		count, _ := strconv.Atoi(fields[1])
		from, _ := strconv.Atoi(fields[3])
		from--
		to, _ := strconv.Atoi(fields[5])
		to--

		for i := 0; i < count; i++ {
			crate := stacks[from].Front()
			stacks[from].Remove(crate)
			stacks[to].PushFront(crate.Value)
		}
	}

	result := ""
	for _, s := range stacks {
		fmt.Printf("s: %v\n", s)

		result += s.Front().Value.(string)
		/*
			for e := s.Front(); e != nil; e = e.Next() {
				fmt.Printf("e: %v\n", e.Value)
			}
		*/
	}
	fmt.Println(result)
}
