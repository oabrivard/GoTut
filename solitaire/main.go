/*
Copyright 2022 Olivier Abrivard

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.
*/

package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/exp/slices"
)

// Create alias to long function names
var pl = fmt.Println

type slot byte

const (
	Empty slot = iota
	HasToken
	Unused
)

type state [][]slot

const BOARD_SIZE = 7

type direction byte

const (
	Down direction = iota
	Left
	Up
	Right
	Invalid
)

func (d direction) isValid() bool {
	return d < Invalid
}

func (d direction) next() direction {
	return (d + 1) % 5
}

func (d direction) toString() string {
	switch d {
	case Down:
		return "down"
	case Left:
		return "left"
	case Up:
		return "up"
	case Right:
		return "right"
	}

	panic("Invalid direction")
}

type tokenPos struct {
	row int
	col int
}

type move struct {
	tokenPos  tokenPos
	direction direction
}

// isEndState verify if the game state s is the final goal state.
// It returns true if the state is the goal state, false otherwise.
func isGoalState(s state) bool {
	goal := state{
		{Unused, Unused, Empty, Empty, Empty, Unused, Unused},
		{Unused, Unused, Empty, Empty, Empty, Unused, Unused},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, HasToken, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Unused, Unused, Empty, Empty, Empty, Unused, Unused},
		{Unused, Unused, Empty, Empty, Empty, Unused, Unused},
	}

	return slices.EqualFunc(s, goal, func(row1, row2 []slot) bool { return slices.Equal(row1, row2) })
}

// getAdjacentTokenPos gets the two token position adjacent
// to token at position pos in direction d.
// It returns the coordinates (tokenPos) of the two adjacent tokens.
func getAdjacentTokenPos(pos tokenPos, d direction) (tokenPos, tokenPos) {
	switch d {
	case Down:
		return tokenPos{pos.row + 1, pos.col}, tokenPos{pos.row + 2, pos.col}
	case Left:
		return tokenPos{pos.row, pos.col - 1}, tokenPos{pos.row, pos.col - 2}
	case Up:
		return tokenPos{pos.row - 1, pos.col}, tokenPos{pos.row - 2, pos.col}
	case Right:
		return tokenPos{pos.row, pos.col + 1}, tokenPos{pos.row, pos.col + 2}
	}

	panic("Invalid direction")
}

// moveToken changes the state s of the board by making the move m.
// The token at position (m.tokenPos.row, m.tokenPos.column) will be removed.
// The adjacent token will also be removed.
// The next adjacent slot will be filled by a token.
func moveToken(s state, m move) {
	adjacent, nextToAdjacent := getAdjacentTokenPos(m.tokenPos, m.direction)

	s[m.tokenPos.row][m.tokenPos.col] = Empty
	s[adjacent.row][adjacent.col] = Empty
	s[nextToAdjacent.row][nextToAdjacent.col] = HasToken
}

// undoMove changes the state s of the board by undoing the move m.
// The slot at position (move.tokenPos.row, move.tokenPos.column) will be filled by a token.
// The adjacent slot will also be filled.
// The next adjacent slot will has its token removed.
func undoMove(s state, m move) {
	adjacent, nextToAdjacent := getAdjacentTokenPos(m.tokenPos, m.direction)

	s[m.tokenPos.row][m.tokenPos.col] = HasToken
	s[adjacent.row][adjacent.col] = HasToken
	s[nextToAdjacent.row][nextToAdjacent.col] = Empty
}

// isMoveAllowed checks if a move is a valid move.
// A move is valid if there exists a token at the location (tokenPos.row, move.tokenPos.column)
// and a token exists at the adjacent slot in the given direction and the is an empty slot just after.
// It returns true if the move is allowed, false otherwise.
func isMoveAllowed(s state, m move) bool {

	adjacent, nextToAdjacent := getAdjacentTokenPos(m.tokenPos, m.direction)

	return nextToAdjacent.col >= 0 && nextToAdjacent.col < BOARD_SIZE &&
		nextToAdjacent.row >= 0 && nextToAdjacent.row < BOARD_SIZE &&
		s[m.tokenPos.row][m.tokenPos.col] == HasToken &&
		s[adjacent.row][adjacent.col] == HasToken &&
		s[nextToAdjacent.row][nextToAdjacent.col] == Empty
}

// depthFirstSearch tries to find a solution to the game using a depth first search approach.
// Current state and already made moves are given as input.
// state will be updated by each move and will be in the "goal state" if we find a solution.
// It returns a slice of valid moves if a solution is found, an empty slice otherwise.
func depthFirstSearch(state state, moves []move) []move {

	// The first way to end the recursive call of this function is if we have found the solution.
	if isGoalState(state) {
		return moves
	}

	// Explore all possible moves, starting by the top left slot.
	for r := range state {
		for c := range state[r] {
			for d := Down; d.isValid(); d = d.next() {
				m := move{
					tokenPos{r, c},
					d,
				}

				// since it is a brute force algorithm, we must discard invalid moves
				if isMoveAllowed(state, m) {
					moveToken(state, m) // updates the state to reflect the move
					moves = append(moves, m)

					// recursive call to find the next move that should lead to the solution
					result := depthFirstSearch(state, moves)

					if len(result) > 0 {
						// we just pop out of the recursive call, returning the solution
						return result
					} else {
						// the move we tried led to a dead end. We must undo it
						moves = moves[:len(moves)-1]
						undoMove(state, m) // updates the state to cancel the move
					}
				}
			}
		}
	}

	return make([]move, 0)
}

// printSolution prints each move from moves that leads to the Solitaire solution
func printSolution(moves []move) {
	if len(moves) > 0 {
		for _, m := range moves {
			_, nextToAdjacent := getAdjacentTokenPos(m.tokenPos, m.direction)

			fmt.Printf("Move %s from (%d,%d) to (%d,%d)\n",
				m.direction.toString(),
				m.tokenPos.row, m.tokenPos.col,
				nextToAdjacent.row, nextToAdjacent.col)
		}
	} else {
		pl("No solution.")
	}
}

// main is the program entry point
func main() {
	start := time.Now()

	// Uncomment the following block to generate a pprof file
	// and profile the code performance
	/*
		f, err := os.Create("profile_solitaire.pb.gz")
		if err != nil {
			log.Fatal(err)
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	*/

	startState := state{
		{Unused, Unused, HasToken, HasToken, HasToken, Unused, Unused},
		{Unused, Unused, HasToken, HasToken, HasToken, Unused, Unused},
		{HasToken, HasToken, HasToken, HasToken, HasToken, HasToken, HasToken},
		{HasToken, HasToken, HasToken, Empty, HasToken, HasToken, HasToken},
		{HasToken, HasToken, HasToken, HasToken, HasToken, HasToken, HasToken},
		{Unused, Unused, HasToken, HasToken, HasToken, Unused, Unused},
		{Unused, Unused, HasToken, HasToken, HasToken, Unused, Unused},
	}

	var moves []move
	moves = depthFirstSearch(startState, moves)

	args := os.Args
	if len(args) == 2 && args[1] == "print" {
		printSolution(moves)
	} else {
		pl("Finished with ", len(moves), " moves.")
	}

	elapsed := time.Since(start)
	pl("It took : ", elapsed)
}
