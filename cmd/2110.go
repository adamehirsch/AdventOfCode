package cmd

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2110Cmd = &cobra.Command{
	Use:                   "day2110",
	Short:                 "2021 Advent of Code Day 6",
	DisableFlagsInUseLine: true,
	Run:                   day2110Func,
}

func init() {
	rootCmd.AddCommand(day2110Cmd)
}

var opens = map[string]string{"[": "]", "(": ")", "{": "}", "<": ">"}
var closes = map[string]int{"]": 57, ")": 3, "}": 1197, ">": 25137}
var completions = map[string]int{")": 1, "]": 2, "}": 3, ">": 4}

type Line []string

func (l Line) String() string {
	f := ""
	for _, v := range l {
		f += v
	}
	return f
}

func getSyntaxChars(f string) [][]string {
	allStrings := [][]string{}
	file, err := utils.Opener(f, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, v := range strings.Split(file, "\n") {
		allStrings = append(allStrings, strings.Split(v, ""))
	}
	return allStrings

}

func findFirstProblem(t Line) (string, int, []string) {
	counter := []string{}
	for _, v := range t {
		nextClose, isOpen := opens[v]
		if isOpen {
			counter = append(counter, nextClose)
			continue
		}

		wrongCharPoints, isClose := closes[v]
		if isClose {
			// if the closing character is the top on the stack, pop it off
			if v == counter[len(counter)-1] {
				counter = counter[:len(counter)-1]
				continue
			} else {
				return v, wrongCharPoints, []string{}
			}
		}
	}

	return "", 0, counter
}

func day2110Func(cmd *cobra.Command, args []string) {
	allChars := getSyntaxChars("data/2110.txt")
	part1Pts := 0
	part2Pts := []int{}
	part2Winner := 0

	for _, row := range allChars {
		_, points, remaining := findFirstProblem(row)
		part1Pts += points

		// incomplete rows will have no points but need scoring
		if points == 0 {
			ts := 0
			for _, v := range utils.Reverse(remaining) {
				ts = (ts * 5) + completions[v]
			}
			part2Pts = append(part2Pts, ts)
		}
	}

	sort.Ints(part2Pts)
	middleVal := int(math.Round(float64(len(part2Pts) / 2)))
	part2Winner = part2Pts[middleVal]

	fmt.Println("Part 1:", part1Pts)
	fmt.Println("Part 2:", part2Winner)

}
