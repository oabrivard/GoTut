package main

import (
	"bufio"
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

type node struct {
	isLeaf   bool
	value    int
	children []node
}

func (n node) String() string {
	if n.isLeaf {
		return fmt.Sprintf("%v", n.value)

	} else {
		return fmt.Sprintf("%v", n.children)
	}
}

func parse(s string) ([]node, int) {
	//fmt.Println("Parsing", s)

	nodes := make([]node, 0)
	curVal := ""

	i := 0
	for i < len(s) {
		switch s[i] {
		case '[':
			//fmt.Println("Start parsing at", i, "being char", string(s[i+1]))
			c, j := parse(s[i+1:])
			n := node{
				isLeaf:   false,
				children: c,
			}
			nodes = append(nodes, n)
			i += j + 1
			//fmt.Println("Continue parsing at", i+1, "being char", string(s[i+1]))
		case ']':
			if curVal != "" {
				v, err := strconv.Atoi(curVal)
				check(err)
				n := node{
					isLeaf: true,
					value:  v,
				}
				nodes = append(nodes, n)
			}
			//fmt.Println("Parsed nodes", nodes, "leaving char at", i)
			return nodes, i
		case ',':
			if curVal != "" {
				v, err := strconv.Atoi(curVal)
				check(err)
				n := node{
					isLeaf: true,
					value:  v,
				}
				nodes = append(nodes, n)
				curVal = ""
			}
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			curVal += string(s[i])
		default:
			fmt.Println(s[i])
			panic("Unexpected char")
		}
		i++
	}

	return nodes, 0
}

/*
func isOrdered(t1 node, t2 node) bool {
	//fmt.Println("-Compare", t1, "to", t2)

	switch {
	case t1.isLeaf && t2.isLeaf:
		fmt.Println("1Compare", t1.value, "vs", t2.value)
		return t1.value <= t2.value
	case t1.isLeaf && !t2.isLeaf:
		fmt.Println("2Compare", t1.value, "vs", t2.children)
		t := node{
			false,
			0,
			make([]node, 0),
		}
		t.children = append(t.children, t1)
		return isOrdered(t, t2)
	case !t1.isLeaf && t2.isLeaf:
		fmt.Println("3Compare", t1.children, "vs", t2.value)
		t := node{
			false,
			0,
			make([]node, 0),
		}
		t.children = append(t.children, t2)
		return isOrdered(t1, t)
	case !t1.isLeaf && !t2.isLeaf:
		fmt.Println("4Compare", t1.children, "vs", t2.children)
		c1 := t1.children
		c2 := t2.children
		for i := 0; i < len(c1) && i < len(c2); i++ {
			if !isOrdered(c1[i], c2[i]) {
				return false
			}
		}

		if len(c2) < len(c1) {
			return false
		} else {
			return true
		}
	}

	panic("Should not happen!")
}
*/
/*
func isOrdered(t1 node, t2 node) bool {
	//fmt.Println("-Compare", t1, "to", t2)

	switch {
	case t1.isLeaf && t2.isLeaf:
		fmt.Println("Compare", t1.value, "vs", t2.value)
		return t1.value <= t2.value
	case t1.isLeaf && !t2.isLeaf:
		fmt.Println("Compare", t1.value, "vs", t2.children)
		if len(t2.children) == 0 {
			return false
		}
		return t1.value <= t2.children[0].value
	case !t1.isLeaf && t2.isLeaf:
		fmt.Println("Compare", t1.children, "vs", t2.value)
		if len(t1.children) == 0 {
			return true
		}
		return t1.children[0].value <= t2.value
	case !t1.isLeaf && !t2.isLeaf:
		fmt.Println("Compare", t1.children, "vs", t2.children)
		c1 := t1.children
		c2 := t2.children
		for i := 0; i < len(c1) && i < len(c2); i++ {
			if !isOrdered(c1[i], c2[i]) {
				return false
			}
		}

		if len(c2) < len(c1) {
			return false
		} else {
			return true
		}
	}

	panic("Should not happen!")
}
*/

func isOrdered(s string, t1 node, t2 node) int {
	//fmt.Println("-Compare", t1, "to", t2)

	switch {
	case t1.isLeaf && t2.isLeaf:
		fmt.Println(s+"- Compare1", t1.value, "vs", t2.value)
		return t2.value - t1.value
	case t1.isLeaf && !t2.isLeaf:
		fmt.Println(s+"- Compare2", t1.value, "vs", t2.children)
		t := node{
			false,
			0,
			make([]node, 0),
		}
		t.children = append(t.children, t1)
		return isOrdered(s+" ", t, t2)
	case !t1.isLeaf && t2.isLeaf:
		fmt.Println(s+"- Compare3", t1.children, "vs", t2.value)
		t := node{
			false,
			0,
			make([]node, 0),
		}
		t.children = append(t.children, t2)
		return isOrdered(s+" ", t1, t)
	case !t1.isLeaf && !t2.isLeaf:
		fmt.Println(s+"- Compare4", t1.children, "vs", t2.children)
		c1 := t1.children
		c2 := t2.children
		cmp := 0
		for i := 0; i < len(c1) && i < len(c2); i++ {
			cmp = isOrdered(s+" ", c1[i], c2[i])
			if cmp == 0 {
				//fmt.Println("Equality")
				continue
			} else {
				return cmp
			}
		}

		if len(c2) < len(c1) {
			fmt.Println("Right side ran out of item")
			return -1
		} else if len(c2) == len(c1) {
			return 0
		} else {
			fmt.Println("Left side ran out of item")
			return 1
		}
	}

	panic("Should not happen!")
}

func main() {
	fileName := os.Args[1]

	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	trees := make([]node, 0)

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		if len(line) == 0 {
			continue
		}
		nodes, _ := parse(strings.TrimPrefix(line, "["))
		fmt.Println("Parsed", line, "to", nodes)
		t := node{
			false,
			0,
			nodes,
		}
		trees = append(trees, t)
	}

	sum := 0
	for i := 0; i < len(trees)/2; i++ {
		fmt.Println("--- Pair ", i+1)
		cmp := isOrdered("", trees[2*i], trees[2*i+1])
		if cmp > 0 {
			sum += i + 1
			fmt.Println("Pair", i+1, "is ordered")
		} else if cmp < 0 {
			fmt.Println("Pair", i+1, "is not ordered")
		} else {
			fmt.Println("Pair", i+1, "is equal")
			panic("Should not happen!")
		}
	}

	fmt.Println("Finished:", sum)
}
