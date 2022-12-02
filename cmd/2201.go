package cmd

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2201Cmd = &cobra.Command{
	Use:                   "day2201",
	Short:                 "2022 Advent of Code Day 01",
	DisableFlagsInUseLine: true,
	Run:                   day2201Func,
}

func init() {
	rootCmd.AddCommand(day2201Cmd)
}

func day2201Func(cmd *cobra.Command, args []string) {
	scanner, err := utils.FileScanner("data/2201.txt", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	currentElf := 0
	elves := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			currentElf++
		} else {
			cals, _ := strconv.Atoi(line)
			elves[currentElf] += cals
		}
	}

	// sorting requires a slice, but there's gotta be a smarter implementation
	elfNumbers := make([]int, len(elves)+1)
	for i := range elfNumbers {
		elfNumbers[i] = i
	}

	sort.Slice(elfNumbers, func(i, j int) bool {
		return elves[elfNumbers[i]] > elves[elfNumbers[j]]
	})

	fmt.Printf("Maximum Single Elf Calories: %d\n", elves[elfNumbers[0]])
	fmt.Printf("Top Three Elves Calories: %d\n", elves[elfNumbers[0]]+elves[elfNumbers[1]]+elves[elfNumbers[2]])

}
