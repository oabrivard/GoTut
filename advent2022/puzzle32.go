package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type valve struct {
	id             string
	rate           int
	adjacentValves []*valve
	costToSibling  map[string]int
}

type bfsStep struct {
	valve *valve
	cost  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const MAX_MINUTES = 26

func computeExpurgedPressure(valves []*valve, opened map[string]int) int {
	total := 0

	for _, v := range valves {

		if opened[v.id] > 0 && opened[v.id] <= MAX_MINUTES {
			total += (MAX_MINUTES - opened[v.id]) * v.rate
		}

	}

	return total
}

func printSteps(valves []*valve, opened map[string]int, maximum int) {
	for _, v := range valves {
		if opened[v.id] > 0 && opened[v.id] <= MAX_MINUTES {
			fmt.Print(v.id, " opened at ", opened[v.id], " at rate ", v.rate, " # ")
		}
	}

	fmt.Print("leads to maximum ", maximum)
	fmt.Print("\n")
}

var maximum = math.MinInt

/*
1  procedure BFS(G, root) is
2      let Q be a queue
3      label root as explored
4      Q.enqueue(root)
5      while Q is not empty do
6          v := Q.dequeue()
7          if v is the goal then
8              return v
9          for all edges from v to w in G.adjacentEdges(v) do

10              if w is not labeled as explored then
11                  label w as explored
12                  w.parent := v
13                  Q.enqueue(w)
*/
func bfsCostToSibling(start *valve, goal string) int {
	fmt.Println("BFS from", start.id, "to", goal)

	firstStep := &bfsStep{
		valve: start,
		cost:  0,
	}

	queue := list.New()
	queue.PushFront(firstStep)

	visited := make(map[string]bool)
	visited[start.id] = true

	for queue.Len() > 0 {
		back := queue.Back()
		queue.Remove(back)
		step := back.Value.(*bfsStep)

		if step.valve.id == goal {
			fmt.Println("BFS from", start.id, "to", goal, ":", step.cost)
			return step.cost
		}

		for _, v := range step.valve.adjacentValves {
			if !visited[v.id] {
				visited[v.id] = true
				newStep := &bfsStep{
					valve: v,
					cost:  step.cost + 1,
				}
				queue.PushFront(newStep)
			}
		}
	}

	panic("Graph is disconnected")
}

/*
procedure DFS(G, v) is

	label v as discovered
	for all directed edges from v to w that are in G.adjacentEdges(v) do
	    if vertex w is not labeled as discovered then
	        recursively call DFS(G, w)
*/
func depthFirstSearch(elephantSearch bool, start *valve, nzv []*valve, elapsed int, opened map[string]int, aa *valve) int {

	max := computeExpurgedPressure(nzv, opened)

	if !elephantSearch {
		// try all elephants option start from "AA"

		// The elephant can only open valves that we haven't open yet.
		nonOpenedNzv := make([]*valve, 0)
		for _, v := range nzv {
			if opened[v.id] == 0 {
				nonOpenedNzv = append(nonOpenedNzv, v)
			}
		}

		openedByElephant := make(map[string]int)

		// Let the elephant run around from "AA" at time zero
		elephantMaxGaz := depthFirstSearch(true, aa, nonOpenedNzv, 0, openedByElephant, aa)

		// This maximum represents the case when we do not open any more
		// valves but we let the elephant run around.
		max += elephantMaxGaz
	}

	for _, v := range nzv {

		if opened[v.id] == 0 {

			shouldOpenAt := elapsed + start.costToSibling[v.id] + 1

			if shouldOpenAt <= MAX_MINUTES {

				opened[v.id] = shouldOpenAt

				currentMaxGaz := depthFirstSearch(elephantSearch, v, nzv, shouldOpenAt, opened, aa)

				if currentMaxGaz > max {
					max = currentMaxGaz
				}

				if max > maximum {
					maximum = max
					printSteps(nzv, opened, maximum)
				}

				opened[v.id] = 0
			}

		}
	}

	return max
}

func main() {
	fileName := os.Args[1]

	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	valves := map[string]*valve{}

	// Parse valves
	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())

		fields := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' })

		s := strings.TrimPrefix(fields[4], "rate=")
		s = strings.TrimSuffix(s, ";")
		r, err := strconv.Atoi(s)
		check(err)

		var v *valve
		var prs bool

		if v, prs = valves[fields[1]]; !prs {
			v = &valve{
				id: fields[1],
			}

			valves[v.id] = v
		}

		v.rate = r
		v.adjacentValves = make([]*valve, 0)
		v.costToSibling = make(map[string]int)

		for i := 9; i < len(fields); i++ {
			id := strings.TrimSuffix(fields[i], ",")

			var av *valve

			if av, prs = valves[id]; !prs {
				av = &valve{
					id: id,
				}

				valves[av.id] = av
			}

			v.adjacentValves = append(v.adjacentValves, av)
		}

		fmt.Println(v)
	}

	// For each valve, compute cost to other valves
	count := 0
	for _, v1 := range valves {
		for _, v2 := range valves {

			_, containsKey := v1.costToSibling[v2.id]

			if v1.id != v2.id && (v1.id == "AA" || v1.rate > 0 && v2.rate > 0) && !containsKey {
				cost := bfsCostToSibling(v1, v2.id)

				v1.costToSibling[v2.id] = cost
				v2.costToSibling[v1.id] = cost
				count++
			}

		}

	}
	fmt.Println(count)

	nonZeroValves := make([]*valve, 0)
	for _, v := range valves {
		if v.rate > 0 {
			nonZeroValves = append(nonZeroValves, v)
		}
	}

	allNzv := make([]*valve, len(nonZeroValves))
	copy(allNzv, nonZeroValves)

	opened := make(map[string]int)
	depthFirstSearch(false, valves["AA"], nonZeroValves, 0, opened, valves["AA"])
}
