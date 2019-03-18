package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type inningResults struct {
	bases []bool
	runs  int
	hits  int
	outs  int
}

type occupiedBases struct {
	first  bool
	second bool
	third  bool
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(input(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func input(actions string) string {
	action := strings.Split(actions, ",")

	i := 0
	inningState := inningResults{
		hits: 0,
		runs: 0,
		outs: 0,
	}

	inningState.bases = make([]bool, 3)

	for range action {
		inningState := simulation(action[i], inningState)
		i++
		if inningState.outs == 3 {
			return fmt.Sprintf("Bases: %v  Outs: %v  Runs: %v", inningState.bases, inningState.outs, inningState.runs)
		}
	}

	return fmt.Sprintf("Bases: %v  Outs: %v  Runs: %v", inningState.bases, inningState.outs, inningState.runs)
}

func simulation(action string, inningState inningResults) inningResults {
	if action != "out" && ((inningState.outs + 1) < 3) {
		if action == "1b" {
			inningState.hits++
			return updateOccupiedBases(inningState, 1)
		} else if action == "hbp" || action == "bb" || action == "out" {
			return updateOccupiedBases(inningState, 1)
		} else if action == "2b" {
			inningState.hits++
			return updateOccupiedBases(inningState, 2)
		} else if action == "3b" {
			inningState.hits++
			return updateOccupiedBases(inningState, 3)
		} else if action == "hr" {
			inningState.hits++
			return updateOccupiedBases(inningState, 4)
		} else if action == "k" || action == "out" {
			return updateOccupiedBases(inningState, 0)
		}
	}
	inningState.outs = 3
	return inningState
}

func updateOccupiedBases(state inningResults, basesToAdvance int) inningResults {
	k := 0
	for ; k <= basesToAdvance; k++ {
		if !state.bases[k] {

		}
	}
	state.bases[k] = true

	return state
}
