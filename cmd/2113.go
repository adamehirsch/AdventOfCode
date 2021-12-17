package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
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

func getCoordsAndFolds(f string) ([]Point, []FoldInstruction) {
	file, err := utils.Opener("data/2113-sample.txt", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	points := []Point{}
	folds := []FoldInstruction{}

	for _, v := range strings.Split(file, "\n") {
		s := strings.Split(v, ",")

		if len(s) == 2 {
			x, _ := strconv.Atoi(s[0])
			y, _ := strconv.Atoi(s[1])

			points = append(points, Point{X: x, Y: y})
		} else if strings.HasPrefix(v, "fold along") {
			s = strings.Split(v, "=")
			l, _ := strconv.Atoi(s[1])
			ax := string(s[0][len(s[0])-1])
			folds = append(folds, FoldInstruction{axis: ax, line: l})
		}
	}
	return points, folds
}

func day2113Func(cmd *cobra.Command, args []string) {
	points, folds := getCoordsAndFolds("data/2113-sample.txt")
	fmt.Println(points, folds)
}
