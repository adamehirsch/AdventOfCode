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

func findHighestJoltageByLength(batteries string, desiredLength int) int {
	finalAnswer := 0
	// we're looking for the highest first digit in the
	// string[:len(string)-desiredLength+1], and then capture the idx of the
	// first digit, so that we can then find the highest digit in
	// string[idx:len(string)-desiredLength-2] and so on until we have a number with desiredLength digits

	workingBatteries := batteries
	for len(workingBatteries) > 0 && desiredLength > 0 {
		highest, highestIdx := findLeftmostHighestDigit(workingBatteries[:len(workingBatteries)-desiredLength+1])
		finalAnswer = finalAnswer*10 + highest
		workingBatteries = workingBatteries[highestIdx+1:]
		desiredLength--
	}
	return finalAnswer
}

func findLeftmostHighestDigit(batteries string) (int, int) {
	highest := 0
	highestIdx := -1
	for idx := 0; idx < len(batteries); idx++ {
		digit, _ := strconv.Atoi(string(batteries[idx]))
		if digit > highest {
			highest = digit
			highestIdx = idx
		}
	}
	return highest, highestIdx
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
	fmt.Printf("Total Joltage with 2: %d \n", totalJoltage)

	totalJoltageByLength := 0

	for _, line := range lines {
		batteries := strings.TrimSpace(line)
		totalJoltageByLength += findHighestJoltageByLength(batteries, 12)
	}
	fmt.Printf("Total Joltage with 12: %d \n", totalJoltageByLength)
}
