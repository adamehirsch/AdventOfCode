package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var day2111Cmd = &cobra.Command{
	Use:                   "day2111",
	Short:                 "2021 Advent of Code Day 11",
	DisableFlagsInUseLine: true,
	Run:                   day2111Func,
}

func init() {
	rootCmd.AddCommand(day2111Cmd)
}

func getMap(f string) [][]int {
	var dm [][]int

	content, err := utils.Opener(f, true)
	if err != nil {
		log.Fatal(err)
	}
	// trim trailing newline

	for _, v := range strings.Split(content, "\n") {
		var row []int

		for _, char := range strings.Split(v, "") {
			f, _ := strconv.Atoi(char)
			row = append(row, f)
		}
		dm = append(dm, row)
	}
	return dm

}

type Point struct {
	X int
	Y int
}

type OctoBoard struct {
	board      [][]int
	stepcount  int
	flashed    []Point
	flashcount int
}

func (ob OctoBoard) String() string {
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	rv := ""
	for _, v := range ob.board {
		for _, o := range v {
			switch o {
			case 0:
				rv += yellow(fmt.Sprintf("%d", o))
			case 9:
				rv += red(fmt.Sprintf("%d", o))
			default:
				rv += fmt.Sprintf("%d", o)
			}
		}
		rv += "\n"
	}
	rv += fmt.Sprintf("\nStep Count: %d\n", ob.stepcount)
	rv += fmt.Sprintf("Flash Count: %d\n", ob.flashcount)
	return rv
}

func (ob *OctoBoard) Neighbors(x, y int) []Point {

	// zero value for bool is false
	var ym, xm, yp, xp bool

	if y > 0 {
		ym = true
	}
	if x > 0 {
		xm = true
	}
	if y < len(ob.board)-1 {
		yp = true
	}
	if x < len(ob.board[0])-1 {
		xp = true
	}
	neighbors := []Point{}
	if xm && ym {
		neighbors = append(neighbors, Point{X: x - 1, Y: y - 1})
	}
	if ym {
		neighbors = append(neighbors, Point{X: x, Y: y - 1})
	}
	if xp && ym {
		neighbors = append(neighbors, Point{X: x + 1, Y: y - 1})
	}
	if xp {
		neighbors = append(neighbors, Point{X: x + 1, Y: y})
	}
	if xp && yp {
		neighbors = append(neighbors, Point{X: x + 1, Y: y + 1})
	}
	if yp {
		neighbors = append(neighbors, Point{X: x, Y: y + 1})
	}
	if xm && yp {
		neighbors = append(neighbors, Point{X: x - 1, Y: y + 1})
	}
	if xm {
		neighbors = append(neighbors, Point{X: x - 1, Y: y})
	}

	return neighbors

}

func (ob *OctoBoard) PointValue(x, y int) *int {
	return &ob.board[y][x]
}

func (ob *OctoBoard) IncreaseEnergy() {
	for y := 0; y < len(ob.board); y++ {
		for x := 0; x < len(ob.board[0]); x++ {
			ob.board[y][x]++
		}
	}
}

func (ob *OctoBoard) BumpNeighbors(x, y int) {
	neighbors := ob.Neighbors(x, y)
	for _, n := range neighbors {
		ob.board[n.Y][n.X]++
	}
}

func (ob *OctoBoard) ClearFlashes() {
	for _, p := range ob.flashed {
		ob.flashcount++
		ob.board[p.Y][p.X] = 0
	}
	ob.flashed = []Point{}
}

func (ob *OctoBoard) Flash() bool {

	flashed := false
	for y := 0; y < len(ob.board); y++ {
		for x := 0; x < len(ob.board[y]); x++ {
			// any point greater than 9 that hasn't already flashed; flash
			if ob.board[y][x] > 9 && !Contains(ob.flashed, Point{X: x, Y: y}) {
				// fmt.Printf("FLASHING: %d, %d\n", x, y)
				flashed = true
				// increase neighbors
				ob.BumpNeighbors(x, y)
				ob.flashed = append(ob.flashed, Point{X: x, Y: y})
			}

		}

	}
	return flashed
}

func (ob *OctoBoard) Step() int {
	ob.IncreaseEnergy()
	for ob.Flash() {
	}
	fc := len(ob.flashed)
	ob.ClearFlashes()
	ob.stepcount++
	return fc
}

func Contains(s []Point, p Point) bool {
	for _, item := range s {
		if item.X == p.X && item.Y == p.Y {
			return true
		}
	}
	return false
}

func day2111Func(cmd *cobra.Command, args []string) {
	OctoMap := OctoBoard{
		board:     getMap("data/2111.txt"),
		stepcount: 0,
	}

	// this time, look for the board where all 100 flashed
	f := 0
	for f < 100 {
		f = OctoMap.Step()

	}
	fmt.Print(OctoMap)

}
