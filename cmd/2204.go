package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2204Cmd = &cobra.Command{
	Use:                   "day2204",
	Short:                 "2022 Advent of Code Day 04",
	DisableFlagsInUseLine: true,
	Run:                   day2204Func,
}

func init() {
	rootCmd.AddCommand(day2204Cmd)
}

func isOverlap(elves []string, phase int) bool {

	var e1, e2, f1, f2 int

	fmt.Sscanf(elves[0], "%d-%d", &e1, &e2)
	fmt.Sscanf(elves[1], "%d-%d", &f1, &f2)

	switch phase {
	case 1:
		if e1 <= f1 && e2 >= f2 {
			return true
		}
		if e1 >= f1 && e2 <= f2 {
			return true
		}
		return false
	case 2:
		switch {
		// if any of the edge points overlap, it's an overlap
		case e1 == f1 || e1 == f2 || e2 == f1 || e2 == f2:
			return true

		case e1 < f1 && f1 < e2:
			return true
		case f1 < e2 && e2 < f2:
			return true

		case f1 < e1 && e1 < f2:
			return true
		case e1 < f2 && f2 < e2:
			return true

		}
		return false
	}
	return false

}

func day2204Func(cmd *cobra.Command, args []string) {
	wholeInput, err := utils.Opener("data/2204.txt", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(wholeInput, "\n")

	overlaps := 0
	o2 := 0
	for _, l := range lines {
		elves := strings.Split(l, ",")
		if isOverlap(elves, 1) {
			overlaps++
		}
		if isOverlap(elves, 1) || isOverlap(elves, 2) {
			o2++
		}
	}
	fmt.Println("Phase 1: ", overlaps)
	fmt.Println("Phase 2: ", o2)
}
