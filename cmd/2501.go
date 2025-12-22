package cmd

import (
	"fmt"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2501Cmd = &cobra.Command{
	Use:                   "day2501",
	Short:                 "2025 Advent of Code Day 01",
	DisableFlagsInUseLine: true,
	Run:                   day2501Func,
}

func init() {
	rootCmd.AddCommand(day2501Cmd)
}

func turnDial(d string, currentPosition *int) bool {
	// d will be in format "Lxx" or "Rxx"
	// dial is a circular dial from 0 to 99
	// interpret the string and return the new position
	if d[0] != 'L' && d[0] != 'R' {
		return false
	}
	var move int
	_, err := fmt.Sscanf(d[1:], "%d", &move)
	if err != nil {
		return false
	}
	if d[0] == 'L' {
		*currentPosition = (*currentPosition - move + 100) % 100
	} else {
		*currentPosition = (*currentPosition + move) % 100
	}
	if *currentPosition == 0 {
		return true
	}
	return false
}

func day2501Func(cmd *cobra.Command, args []string) {
	file, err := utils.Opener("data/2501.txt", true)
	if err != nil {
		fmt.Printf("failed to open input: %v\n", err)
		return
	}

	dialPosition := 50
	howManyZeroes := 0

	for _, v := range strings.Split(file, "\n") {
		if turnDial(v, &dialPosition) {
			howManyZeroes++
		}
	}

	fmt.Printf("The dial landed on 0 a total of %d times\n", howManyZeroes)
}
