package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2203Cmd = &cobra.Command{
	Use:                   "day2203",
	Short:                 "2022 Advent of Code Day 03",
	DisableFlagsInUseLine: true,
	Run:                   day2203Func,
}

var alphabetMap map[string]int

func init() {
	rootCmd.AddCommand(day2203Cmd)
	alphabetMap = make(map[string]int)
	for i, letter := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		alphabetMap[string(letter)] = i + 1
	}
}

func findSplitLetters(s string) string {
	l := strings.Split(s, "")
	firstHalf := strings.Join(l[:len(l)/2], "")
	secondHalf := strings.Join(l[len(l)/2:], "")

	if findCommonLetter(firstHalf, secondHalf) != "" {
		return (findCommonLetter(firstHalf, secondHalf))
	}
	return ""
}

func findCommonLetter(s ...string) string {
	seen := make(map[string]int)

	first := s[0]

	// for every letter in "first"
	for _, firstLetter := range first {
		if seen[string(firstLetter)] == 0 {
			seen[string(firstLetter)] = 1
		} else {
			continue
		}
		// loop over each of the remaining sent strings
		for _, otherSet := range s[1:] {
			// for every letter in each string
			if strings.Contains(otherSet, string(firstLetter)) {
				seen[string(firstLetter)] += 1
			}
		}
	}
	for letter, howMany := range seen {
		// if any letter has been seen as many times as the sets handed in
		if howMany >= len(s) {
			return (string(letter))
		}
	}
	return ""
}

func letterValue(s string) int {
	return alphabetMap[s]
}

func day2203Func(cmd *cobra.Command, args []string) {
	wholeInput, err := utils.Opener("data/2203.txt", true)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := strings.Split(wholeInput, "\n")
	total := 0

	for _, line := range lines {
		total += letterValue(findSplitLetters(line))
	}

	fmt.Printf("Phase 1: Total: %d\n", total)

	// Create a slice to hold groups of three lines.
	groups := make([][]string, 0)

	// Loop over the lines and append groups of three to the groups slice.
	for i := 0; i < len(lines); i += 3 {
		group := lines[i : i+3]
		groups = append(groups, group)
	}

	secondTotal := 0
	for _, g := range groups {
		secondTotal += letterValue(findCommonLetter(g...))
	}
	fmt.Printf("Phase 2 total: %d\n", secondTotal)
}
