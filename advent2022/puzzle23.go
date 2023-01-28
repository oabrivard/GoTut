package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

type node struct {
	value      rune
	isExplored bool
	i          int
	j          int
	isGoal     bool
	parent     *node
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func adjacentNodes(graph [][]*node, n *node) []*node {
	result := make([]*node, 0)
	i := n.i
	j := n.j

	if i > 0 {
		result = append(result, graph[i-1][j])
	}

	if i < len(graph)-1 {
		result = append(result, graph[i+1][j])
	}

	if j > 0 {
		result = append(result, graph[i][j-1])
	}

	if j < len(graph[0])-1 {
		result = append(result, graph[i][j+1])
	}

	return result
}

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
func BFS(graph [][]*node, start *node) *node {
	queue := list.New()
	start.isExplored = true
	queue.PushFront(start)

	for queue.Len() > 0 {
		back := queue.Back()
		queue.Remove(back)
		v := back.Value.(*node)

		if v.isGoal {
			return v
		}

		for _, w := range adjacentNodes(graph, v) {
			if !w.isExplored {
				if w.value-v.value <= 1 {
					w.isExplored = true
					w.parent = v
					queue.PushFront(w)
				}
			}
		}
	}

	return nil
}

func main() {
	fileName := os.Args[1]

	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	graph := [][]*node{}
	var start *node

	i := 0
	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		graph = append(graph, make([]*node, len(line)))

		j := 0
		for _, b := range line {
			graph[i][j] = &node{
				value:      b,
				isExplored: false,
				i:          i,
				j:          j,
				isGoal:     false,
				parent:     nil,
			}

			if b == 'S' {
				start = graph[i][j]
				start.value = 'a'
			}

			if b == 'E' {
				graph[i][j].isGoal = true
				graph[i][j].value = 'z'
			}

			j++
		}

		i++
	}

	fmt.Println(graph, start)

	goal := BFS(graph, start)

	if goal == nil {
		fmt.Println("Solution not found")
	} else {
		i := 0
		for n := goal.parent; n != nil; n = n.parent {
			fmt.Print(string(n.value), "->")
			i++
		}
		fmt.Println("\nShortest path is", i)
	}

}
