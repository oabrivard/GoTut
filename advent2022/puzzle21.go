package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	id          int
	items       []int
	operation   func(int) int
	destination func(int) int
	inspections int
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
	var current int
	monkeys := []*monkey{}

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())

		switch {
		case strings.HasPrefix(line, "Monkey "):
			s := strings.TrimPrefix(line, "Monkey ")
			s = strings.TrimSuffix(s, ":")
			id, _ := strconv.Atoi(s)
			current = id
			m := monkey{
				id:    id,
				items: make([]int, 0),
			}
			monkeys = append(monkeys, &m) // appends Monkey 0 to slot 0 since they are ordered in the text file
		case strings.HasPrefix(line, "Starting items: "):
			s := strings.TrimPrefix(line, "Starting items: ")
			items := strings.Split(s, ", ")
			for _, item := range items {
				i, _ := strconv.Atoi(item)
				monkeys[current].items = append(monkeys[current].items, i)
			}
		case strings.HasPrefix(line, "Operation: new = old "):
			s := strings.TrimPrefix(line, "Operation: new = old ")
			fields := strings.Fields(s)
			if fields[0] == "*" {
				if fields[1] == "old" {
					monkeys[current].operation = func(i int) int { return i * i }
				} else {
					val, _ := strconv.Atoi(fields[1])
					monkeys[current].operation = func(i int) int { return i * val }
				}
			} else {
				if fields[1] == "old" {
					monkeys[current].operation = func(i int) int { return i + i }
				} else {
					val, _ := strconv.Atoi(fields[1])
					monkeys[current].operation = func(i int) int { return i + val }
				}
			}
		case strings.HasPrefix(line, "Test: divisible by "):
			s := strings.TrimPrefix(line, "Test: divisible by ")
			divider, _ := strconv.Atoi(s)
			fileScanner.Scan()
			line = strings.TrimSpace(fileScanner.Text())
			s = strings.TrimPrefix(line, "If true: throw to monkey ")
			valTrue, _ := strconv.Atoi(s)
			fileScanner.Scan()
			line = strings.TrimSpace(fileScanner.Text())
			s = strings.TrimPrefix(line, "If false: throw to monkey ")
			valFalse, _ := strconv.Atoi(s)
			monkeys[current].destination = func(i int) int {
				if i%divider == 0 {
					return valTrue
				} else {
					return valFalse
				}
			}
		case line == "":
			fmt.Println(monkeys[current])
		default:
			panic(fmt.Sprintf("invalid instruction '%s'", line))
		}
	}

	fmt.Println(monkeys)

	for z := 0; z < 20; z++ {
		for _, m := range monkeys {
			fmt.Println("Monkey", m.id)

			for len(m.items) > 0 {
				head := m.items[0]
				fmt.Println("inspects", head)
				m.items = m.items[1:]
				m.inspections++
				lvl := m.operation(head)
				lvl = lvl / 3
				dest := m.destination(lvl)
				fmt.Println("sends worry level", lvl, "to dest", dest)
				monkeys[dest].items = append(monkeys[dest].items, lvl)
			}
		}
		fmt.Println("Monkey 0 :", monkeys[0].inspections)
		fmt.Println("Monkey 1 :", monkeys[1].inspections)
		fmt.Println("Monkey 2 :", monkeys[2].inspections)
		fmt.Println("Monkey 3 :", monkeys[3].inspections)
		fmt.Println("----------")
	}

	inspections := []int{}
	for _, m := range monkeys {
		inspections = append(inspections, m.inspections)
	}
	sort.Ints(inspections)
	fmt.Println(inspections)
	businessLvl := inspections[len(inspections)-1] * inspections[len(inspections)-2]
	fmt.Println(businessLvl)
}
