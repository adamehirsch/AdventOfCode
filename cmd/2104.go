package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var day2104Cmd = &cobra.Command{
	Use:                   "day2104",
	Short:                 "2021 Advent of Code Day 4",
	DisableFlagsInUseLine: true,
	Run:                   day2104Func,
}

func init() {
	rootCmd.AddCommand(day2104Cmd)
}

type bingoBoard [][]string

func (bb bingoBoard) String() string {
	out := ""
	for _, row := range bb {
		for _, v := range row {
			out += fmt.Sprintf("%3s", v)
		}
		out += "\n"
	}
	return out
}

func (bb *bingoBoard) printWinner(wn []string) int {

	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	colorizedBB := bingoBoard{}
	unmarkedTotal := 0
	for _, v := range *bb {
		var currRow []string

		for _, val := range v {
			v := ""
			if utils.Contains(wn, val) {
				v = green(fmt.Sprintf("%3s", val))
			} else {
				// add any un-called number to the
				n, _ := strconv.Atoi(val)
				unmarkedTotal += n

				v = red(fmt.Sprintf("%3s", val))
			}
			currRow = append(currRow, v)
		}
		colorizedBB = append(colorizedBB, currRow)
	}
	fmt.Println(colorizedBB)
	return unmarkedTotal
}

func (bb *bingoBoard) isWinner(wn []string) bool {

	// check all the rows
	for i := 0; i < len(*bb); i++ {
		if rowWins((*bb)[i], wn) {
			return true
		}
	}

	// check all the columns
	for i := 0; i < 5; i++ {
		col := make([]string, 0)
		for j := 0; j < 5; j++ {
			col = append(col, (*bb)[j][i])
		}
		if rowWins(col, wn) {
			return true
		}

	}
	return false
}

func rowWins(row []string, wn []string) bool {
	for _, v := range row {
		if !utils.Contains(wn, v) {
			return false
		}
	}
	return true
}

func getBoards() []bingoBoard {
	allBoards := []bingoBoard{}

	file, err := utils.Opener("data/2104.txt", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var currBoard bingoBoard

	for _, v := range strings.Split(file, "\n") {
		b := strings.Fields(v)
		// we've got a good row only if there are 5 entries
		if len(b) == 5 {
			currBoard = append(currBoard, b)
		}

		// if we've completed a 5x5 board, add it to the list
		if len(currBoard) == 5 {
			allBoards = append(allBoards, currBoard)
			currBoard = bingoBoard{}
		}

	}

	return allBoards
}

func getWinningNums() []string {
	allNums := []string{}

	file, err := utils.Opener("data/2104-numbers.txt", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	allNums = append(allNums, strings.Split(file, ",")...)

	return allNums
}

func day2104Func(cmd *cobra.Command, args []string) {
	allBoards := getBoards()
	allNums := getWinningNums()

	var incNums []string
	var firstWinningBoard bingoBoard
	var firstWinningNum int
	var firstWinningIncNums []string

	var lastWinningBoard bingoBoard
	var lastWinningNum int
	var lastWinningIncNums []string

	for _, v := range allNums {
		incNums = append(incNums, v)

		for i := len(allBoards) - 1; i >= 0; i-- {
			b := allBoards[i]
			if b.isWinner(incNums) {
				if len(firstWinningBoard) == 0 {
					firstWinningBoard = b
					firstWinningNum, _ = strconv.Atoi(v)
					firstWinningIncNums = incNums
				}

				lastWinningBoard = b
				lastWinningNum, _ = strconv.Atoi(v)
				lastWinningIncNums = incNums

				// drop boards that have already won out of the running
				allBoards = append(allBoards[:i], allBoards[i+1:]...)

			}
		}

	}
	color.Blue("First Winning Board")
	winningTotal := firstWinningBoard.printWinner(firstWinningIncNums)
	fmt.Println(firstWinningNum)
	fmt.Printf("%d * %d = %d\n", winningTotal, firstWinningNum, winningTotal*firstWinningNum)

	color.Yellow("Last Winnning Board")
	lWinningTotal := lastWinningBoard.printWinner(lastWinningIncNums)
	fmt.Println(lastWinningNum)
	fmt.Printf("%d * %d = %d\n", lWinningTotal, lastWinningNum, lWinningTotal*lastWinningNum)
}
