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

func isOverlap(e []string) bool {

	var a, b, c, d int

	fmt.Sscanf(e[0], "%d-%d", &a, &b)
	fmt.Sscanf(e[1], "%d-%d", &c, &d)

	if a <= c {
		if b >= d {
			return true
		}
	}
	if a >= c {
		if b <= d {
			return true
		}
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
	for _, l := range lines {
		elves := strings.Split(l, ",")
		if isOverlap(elves) {
			overlaps++
		}
	}
	fmt.Println("Phase 1: ", overlaps)
}
