package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2502Cmd = &cobra.Command{
	Use:                   "day2502",
	Short:                 "2025 Advent of Code Day 02",
	DisableFlagsInUseLine: true,
	Run:                   day2502Func,
}

func init() {
	rootCmd.AddCommand(day2502Cmd)
}

func isNumberRepeated01(number int) bool {
	numberLength := len(fmt.Sprintf("%d", number))
	if numberLength%2 == 0 {
		return strconv.Itoa(number)[:numberLength/2] == strconv.Itoa(number)[numberLength/2:]
	}
	return false

}

func isNumberRepeated02(number int) bool {
	wholeStr := strconv.Itoa(number)
	numberLength := len(wholeStr)

	for subLen := 1; subLen < numberLength; subLen++ {
		// is the whole string evenly divisible by subLen?
		if numberLength%subLen != 0 {
			continue
		}
		part := wholeStr[:subLen]
		if isOnlySubstring(wholeStr, part) {
			return true
		}
	}
	return false
}

func isOnlySubstring(whole, part string) bool {
	return float64(strings.Count(whole, part)) == float64(len(whole))/float64(len(part))
}

func day2502Func(cmd *cobra.Command, args []string) {
	file, err := utils.Opener("data/2502.txt", true)
	if err != nil {
		fmt.Printf("failed to open input: %v\n", err)
		return
	}
	allTheRanges := strings.Split(file, ",")
	totalCount := 0
	for _, r := range allTheRanges {
		bounds := strings.Split(r, "-")
		lowerBound, err := strconv.Atoi(bounds[0])
		if err != nil {
			fmt.Printf("failed to convert lower bound %s to int: %v\n", bounds[0], err)
			return
		}
		upperBound, err := strconv.Atoi(bounds[1])
		if err != nil {
			fmt.Printf("failed to convert upper bound %s to int: %v\n", bounds[1], err)
			return
		}
		count := 0
		for i := lowerBound; i <= upperBound; i++ {
			if isNumberRepeated02(i) {
				fmt.Printf("   Number %d is made of repeated segments\n", i)
				count++
				totalCount += i
			}
		}
		fmt.Printf("In range %d-%d, there are %d numbers with repeated halves\n", lowerBound, upperBound, count)
	}
	fmt.Printf("Total value of numbers with repeated halves: %d\n", totalCount)
}
