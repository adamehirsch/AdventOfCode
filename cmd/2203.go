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
	firstHalf := l[:len(l)/2]
	secondHalf := l[len(l)/2:]

	return ""
}

func findCommonLetter(s ...string) string {
	seen := make(map[rune]int)

	first := s[0]

	// for every letter in "first"
	for _, r := range first {
		// loop over each of the remaining sent strings
		for _, c := range s[1:] {
			// for every letter in each string
			if strings.Contains(c, string(r)) {
				seen[r]++
			}
		}
	}

	for k, v := range seen {
		// if any letter has been seen as many times as
		if v >= len(s) {

		}
	}

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
	fmt.Printf("Total: %d\n", total)

}
