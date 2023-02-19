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
	openedAt       int
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

func computeExpurgedPressure(path []*valve) int {
	total := 0

	for _, v := range path {

		if v.openedAt <= MAX_MINUTES && v.openedAt > 0 {
			total += (MAX_MINUTES - v.openedAt) * v.rate
		}

	}

	return total
}

func printSteps(path []*valve, maximum int) {
	for _, v := range path {
		fmt.Print(v.id, " opened at ", v.openedAt, " at rate ", v.rate, " # ")
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
func depthFirstSearch(start1 *valve, start2 *valve, nzv []*valve, elapsed1 int, elapsed2 int, valvesLeft int, path1 []*valve, path2 []*valve, allNzv []*valve) int {

	// !!! J'AI L'IMPRESSSION QUE PATH DEVRAIT ETRE UNE PAIRE
	//fmt.Println("Start =", start.id, "at", elapsed)

	if elapsed1 >= 30 || elapsed2 >= 30 || valvesLeft == 0 {
		return computeExpurgedPressure(allNzv)
	}

	// Explore all possible pairs
	for i := range nzv {

		v1 := nzv[i]
		if v1 != nil {
			nzv[i] = nil
			path1 = append(path1, v1)

			cost1 := start1.costToSibling[v1.id]
			cost1++
			v1.openedAt = elapsed1 + cost1

			for j := range nzv {

				v2 := nzv[j]

				if v2 != nil {

					nzv[j] = nil
					path2 = append(path2, v2)

					cost2 := start2.costToSibling[v2.id]
					cost2++
					v2.openedAt = elapsed2 + cost2

					totalGaz := depthFirstSearch(v1, v2, nzv, v1.openedAt, v2.openedAt, valvesLeft-2, path1, path2, allNzv)

					if totalGaz > maximum {
						maximum = totalGaz
						fmt.Print("Me : ")
						printSteps(path1, maximum)
						fmt.Print("Elephant : ")
						printSteps(path2, maximum)
					}

					v2.openedAt = -1
					path2 = path2[:len(path2)-1]
					nzv[j] = v2
				}
			}

			v1.openedAt = -1
			path1 = path1[:len(path1)-1]
			nzv[i] = v1
		}
	}

	return computeExpurgedPressure(allNzv)
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
		v.openedAt = -1

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

	path1 := make([]*valve, 0)
	path1 = append(path1, valves["AA"])
	path2 := make([]*valve, 0)
	path2 = append(path2, valves["AA"])
	depthFirstSearch(valves["AA"], valves["AA"], nonZeroValves, 0, 0, len(nonZeroValves), path1, path2, allNzv)
}
