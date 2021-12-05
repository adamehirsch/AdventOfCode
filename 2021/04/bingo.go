package main

import "fmt"

type bingoBoard [][]string

var winningNums []string
var allBoards []bingoBoard

func (bb *bingoBoard) isWinner() bool {

	// check all the rows
	for i := 0; i < len(*bb); i++ {
		if rowWins((*bb)[i]) == true {
			return true
		}
	}

	// check all the columns
	for i := 0; i < 5; i++ {
		col := make([]string, 0)
		for j := 0; j < 5; j++ {
			col = append(col, (*bb)[j][i])
		}
		if rowWins(col) == true {
			return true
		}

	}
	return false
}

func rowWins(row []string) bool {
	for _, v := range row {
		if Contains(winningNums, v) == false {
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

func main() {

	sampleBoard := bingoBoard{
		{"4", "95", "84", "51", "36"},
		{"43", "40", "37", "23", "85"},
		{"14", "90", "8", "59", "99"},
		{"0", "88", "68", "93", "81"},
		{"25", "6", "55", "19", "48"}}

	winningNums = []string{"55", "8", "84", "37", "68"}

	fmt.Print(sampleBoard.isWinner())

	// boards, err := os.Open("input.txt")

	// if err != nil {
	// 	log.Fatalf("failed to open")

	// }
	// defer boards.Close()

	// boardScanner := bufio.NewScanner(boards)

	// var currBoard bingoBoard

	// for boardScanner.Scan() {
	// 	b := strings.Fields(boardScanner.Text())

	// 	if len(currBoard) == 5 {
	// 		allBoards = append(allBoards, currBoard)

	// 	}
	// 	currBoard = append(currBoard, b)

	// 	fmt.Println(b)
	// 	fmt.Println(currBoard)
	// }

	// if err := boardScanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

}
