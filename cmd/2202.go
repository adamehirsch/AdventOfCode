package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2202Cmd = &cobra.Command{
	Use:                   "day2202",
	Short:                 "2022 Advent of Code Day 01",
	DisableFlagsInUseLine: true,
	Run:                   day2202Func,
}

func init() {
	rootCmd.AddCommand(day2202Cmd)
}

func findOutcome(o, m string) int {

	const xPts = 1
	const yPts = 2
	const zPts = 3
	const winPoints = 6
	const losePoints = 0
	const drawPoints = 3

	score := 0
	switch m {
	case "X":
		score += xPts
	case "Y":
		score += yPts
	case "Z":
		score += zPts
	}

	switch o {
	case "A": // rock
		switch m {
		case "X": // Rock
			score += drawPoints
		case "Y": // Paper
			score += winPoints
		case "Z": // scissors
			score += losePoints
		}
	case "B": // paper
		switch m {
		case "X": // Rock
			score += losePoints
		case "Y": // Paper
			score += drawPoints
		case "Z": // scissors
			score += winPoints
		}
	case "C": // scissors
		switch m {
		case "X": // Rock
			score += winPoints
		case "Y": // Paper
			score += losePoints
		case "Z": // scissors
			score += drawPoints
		}
	}
	return score
}

func scoreRound(o, m string, r int) int {

	if r == 1 {
		return findOutcome(o, m)
	}
	// part 2 challenge; win/lose/draw
	wld := map[string][]string{
		"A": {"Y", "Z", "X"},
		"B": {"Z", "X", "Y"},
		"C": {"X", "Y", "Z"},
	}

	ourPlay := ""
	switch m {
	case "X": // we need to lose
		ourPlay = wld[o][1]
	case "Y": // we need to draw
		ourPlay = wld[o][2]
	case "Z": // we need to win
		ourPlay = wld[o][0]
	}

	return findOutcome(o, ourPlay)
}

func day2202Func(cmd *cobra.Command, args []string) {
	scanner, err := utils.FileScanner("data/2202.txt", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	roundOne := 0
	roundTwo := 0
	for scanner.Scan() {
		playedMoves := strings.Fields(scanner.Text())
		roundOne += scoreRound(playedMoves[0], playedMoves[1], 1)
		roundTwo += scoreRound(playedMoves[0], playedMoves[1], 2)
	}

	fmt.Printf("Round 1: %d\n", roundOne)
	fmt.Printf("Round 2: %d\n", roundTwo)
}
