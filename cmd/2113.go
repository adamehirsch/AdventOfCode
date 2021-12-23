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

var day2113Cmd = &cobra.Command{
	Use:                   "day2113",
	Short:                 "2021 Advent of Code Day 13",
	DisableFlagsInUseLine: true,
	Run:                   day2113Func,
}

func init() {
	rootCmd.AddCommand(day2113Cmd)
}

type FoldInstruction struct {
	axis string
	line int
}

func (fi FoldInstruction) String() string {
	return fmt.Sprintf("{ axis: %s, line: %d }", fi.axis, fi.line)
}

func getCoordsAndFolds(f string) ([]Point, []FoldInstruction, int, int) {
	file, err := utils.Opener(f, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Points := []Point{}
	folds := []FoldInstruction{}

	maxX := 0
	maxY := 0

	for _, v := range strings.Split(file, "\n") {
		s := strings.Split(v, ",")

		if len(s) == 2 {
			x, _ := strconv.Atoi(s[0])
			y, _ := strconv.Atoi(s[1])

			if x > maxX {
				maxX = x
			}

			if y > maxY {
				maxY = y
			}

			Points = append(Points, Point{X: x, Y: y})
		} else if strings.HasPrefix(v, "fold along") {
			s = strings.Split(v, "=")
			l, _ := strconv.Atoi(s[1])
			ax := string(s[0][len(s[0])-1])
			folds = append(folds, FoldInstruction{axis: ax, line: l})
		}
	}
	return Points, folds, maxX, maxY
}

func CreateGrid(x, y int) grid {
	fmt.Printf("MAKING X: %d Y: %d\n", x, y)
	// arguments are MaxPosition rather than length
	g := make(grid, y+1)
	for f := 0; f <= y; f++ {
		g[f] = make([]int, x+1)
	}
	return g
}

func PlotPoint(g grid, p Point) {
	// all slices are pointer objects, so you don't need to pointer to them
	g[p.Y][p.X] = 1
}

func PrintGrid(g grid) {
	red := color.New(color.FgYellow).SprintFunc()

	fmt.Printf("PrintGrid: X: %d, Y: %d\n", g.MaxPosX(), g.MaxPosY())
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			v := ""
			if g[y][x] > 0 {
				v = red(fmt.Sprintf("%2s", "#"))
			} else {
				v = "."
			}
			fmt.Printf("%2s", v)
		}
		fmt.Print("\n")
	}
}

func (g grid) MaxPosX() int {
	return len(g[0]) - 1
}

func (g grid) MaxPosY() int {
	return len(g) - 1
}

func FoldGrid(og grid, fi FoldInstruction) grid {
	var ng grid
	fmt.Println("FOLDING: ", fi)

	if fi.axis == "x" {
		// ng = CreateGrid(og.MaxPosX()-(fi.line+1), og.MaxPosY())
		ng = CreateGrid(fi.line, og.MaxPosY())

	} else {
		// ng = CreateGrid(og.MaxPosX(), og.MaxPosY()-(fi.line+1))
		ng = CreateGrid(og.MaxPosX(), fi.line)
	}

	for y, row := range og {
		for x, v := range row {
			if v > 0 {
				if y <= ng.MaxPosY() && x <= ng.MaxPosX() {
					// points on the new grid that simply are identical to those on the OG
					ng[y][x] = v
				} else if y > fi.line && fi.axis == "y" {
					// this Y point exists on og but not on ng
					distanceFromFold := y - fi.line
					// fmt.Printf("y: %d - %d = %d\n", fi.line, distanceFromFold, fi.line-distanceFromFold)
					ng[(fi.line)-distanceFromFold][x] = v

				} else if x > fi.line && fi.axis == "x" {
					// this X point exists on og but not on ng
					distanceFromFold := x - fi.line
					ng[y][(fi.line)-distanceFromFold] = v
				}
			}
		}
	}

	return ng
}

func CountPoints(g grid) int {
	a := 0
	for _, row := range g {
		for _, v := range row {
			if v > 0 {
				a = a + 1
			}
		}
	}
	return a
}

func day2113Func(cmd *cobra.Command, args []string) {
	points, folds, maxX, maxY := getCoordsAndFolds("data/2113.txt")

	grid := CreateGrid(maxX, maxY)
	for _, Point := range points {
		PlotPoint(grid, Point)
	}

	for _, f := range folds {
		grid = FoldGrid(grid, f)
		fmt.Println("Points: ", CountPoints(grid))

	}

	PrintGrid(grid)
	// fmt.Println(CountPoints(grid))
}
