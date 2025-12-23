package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2503Cmd = &cobra.Command{
	Use:                   "day2503",
	Short:                 "2025 Advent of Code Day 03",
	DisableFlagsInUseLine: true,
	Run:                   day2503Func,
}

func init() {
	rootCmd.AddCommand(day2503Cmd)
}

func findHighestJoltage(batteries string) int {
	highest := 0
	nextHighest := 0
	for idx := 0; idx < len(batteries)-1; idx++ {
		digit, _ := strconv.Atoi(string(batteries[idx]))
		if digit > highest {
			highest = digit
			nextHighest = 0
		} else if digit > nextHighest {
			nextHighest = digit
		}
	}

	lastDigit, _ := strconv.Atoi(string(batteries[len(batteries)-1]))
	if lastDigit > nextHighest {
		nextHighest = lastDigit
	}
	return highest*10 + nextHighest

}

func day2503Func(cmd *cobra.Command, args []string) {
	lines, err := utils.OpenerToLines("data/2503.txt", true)
	if err != nil {
		fmt.Printf("failed to open input: %v\n", err)
		return
	}
	totalJoltage := 0

	for _, line := range lines {
		batteries := strings.TrimSpace(line)
		totalJoltage += findHighestJoltage(batteries)
	}
	fmt.Printf("Total Joltage: %d \n", totalJoltage)
}
