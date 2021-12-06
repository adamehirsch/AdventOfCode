package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

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
			if Contains(wn, val) {
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
		if rowWins((*bb)[i], wn) == true {
			return true
		}
	}

	// check all the columns
	for i := 0; i < 5; i++ {
		col := make([]string, 0)
		for j := 0; j < 5; j++ {
			col = append(col, (*bb)[j][i])
		}
		if rowWins(col, wn) == true {
			return true
		}

	}
	return false
}

func rowWins(row []string, wn []string) bool {
	for _, v := range row {
		if Contains(wn, v) == false {
			return false
		}
	}
	return true
}

func Contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func getBoards() []bingoBoard {
	allBoards := []bingoBoard{}

	boards, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer boards.Close()

	boardScanner := bufio.NewScanner(boards)

	var currBoard bingoBoard

	for boardScanner.Scan() {
		b := strings.Fields(boardScanner.Text())

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

	if err := boardScanner.Err(); err != nil {
		log.Fatal(err)
	}
	return allBoards
}

func getWinningNums() []string {
	allNums := []string{}

	nums, err := os.Open("numbers.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer nums.Close()

	numScanner := bufio.NewScanner(nums)

	for numScanner.Scan() {
		for _, v := range strings.Split(numScanner.Text(), ",") {
			allNums = append(allNums, v)
		}

	}

	if err := numScanner.Err(); err != nil {
		log.Fatal(err)
	}
	return allNums
}

func main() {
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
