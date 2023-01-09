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

	dirs := list.New()
	sizes := map[string]int{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fields := strings.Fields((line))

		if fields[0] == "$" && fields[1] == "cd" {
			// a cd command
			fmt.Println("cd", fields[2])

			if fields[2] == ".." {
				curdir := dirs.Front().Value.(string)
				dirs.Remove(dirs.Front())
				parent := dirs.Front().Value.(string)
				sizes[parent] += sizes[curdir]
			} else {
				if dirs.Len() == 0 {
					dirs.PushFront(fields[2])
				} else {
					parent := dirs.Front().Value.(string)
					dirs.PushFront(parent + "/" + fields[2])
				}
			}
		} else if fields[0] != "$" && fields[0] != "dir" {
			fmt.Println("file", fields[1], fields[0])
			// a file
			size, _ := strconv.Atoi(fields[0])
			curdir := dirs.Front().Value.(string)
			sizes[curdir] += size
		}
	}

	for dirs.Len() > 1 {
		curdir := dirs.Front().Value.(string)
		dirs.Remove(dirs.Front())
		parent := dirs.Front().Value.(string)
		sizes[parent] += sizes[curdir]
	}

	fmt.Println(sizes)

	used := sizes["/"]
	unused := 70000000 - used
	needed := 30000000 - unused
	fmt.Println(needed)

	minsize := 70000000
	for _, v := range sizes {
		if v >= needed && v < minsize {
			minsize = v
		}
	}
	fmt.Println(minsize)
}
