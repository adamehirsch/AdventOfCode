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

	calorieInventory := make(map[int]int)
	elfNames := []int{0}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// New elf, new elf "name"
			elfNames = append(elfNames, elfNames[len(elfNames)-1]+1)
		} else {
			cals, _ := strconv.Atoi(line)
			calorieInventory[elfNames[len(elfNames)-1]] += cals
		}
	}

	sort.Slice(elfNames, func(i, j int) bool {
		return calorieInventory[elfNames[i]] > calorieInventory[elfNames[j]]
	})

	fmt.Printf("Maximum Single Elf Calories: %d\n", calorieInventory[elfNames[0]])
	fmt.Printf("Top Three Elves Calories: %d\n", calorieInventory[elfNames[0]]+calorieInventory[elfNames[1]]+calorieInventory[elfNames[2]])

}
