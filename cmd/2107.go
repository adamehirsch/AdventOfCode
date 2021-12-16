package cmd

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2107Cmd = &cobra.Command{
	Use:                   "day2107",
	Short:                 "2021 Advent of Code Day 7",
	DisableFlagsInUseLine: true,
	Run:                   day2107Func,
}

func init() {
	rootCmd.AddCommand(day2107Cmd)
}

func getCrabs(f string) []int {
	file, err := utils.Opener("data/2107.txt", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// initialize an empty map
	var allCrabs []int

	// populate the slice with the crabs
	for _, v := range strings.Split(file, ",") {
		f, _ := strconv.Atoi(v)
		allCrabs = append(allCrabs, f)
	}
	return allCrabs
}

func GetMax(s *[]int) int {
	max := (*s)[1]
	for _, v := range *s {
		if v > max {
			max = v
		}
	}
	return max
}
func GetMin(s *[]int) int {
	min := (*s)[1]
	for _, v := range *s {
		if v < min {
			min = v
		}
	}
	return min
}

func CalcFuelCost(d int) int {
	fc := 0
	for i := 0; i <= d; i++ {
		fc += i
	}
	return fc
}

func Abs(x, y int) int {
	z := x - y
	if z < 0 {
		z = -z
	}
	return z
}

func day2107Func(cmd *cobra.Command, args []string) {
	allCrabs := getCrabs("data/2107.txt")
	sort.Ints(allCrabs)
	day1Costs := map[int]int{}
	day2Costs := map[int]int{}

	// Fuel costs 1 per space
	for j := GetMin(&allCrabs); j <= GetMax(&allCrabs); j++ {
		for _, v := range allCrabs {
			day1Costs[j] += Abs(v, j)
			day2Costs[j] += CalcFuelCost(Abs(v, j))
		}
	}

	var fc1, fc2 int
	// find Part 1
	for _, v := range day1Costs {
		if fc1 == 0 {
			fc1 = v
		}
		if v < fc1 {
			fc1 = v
		}

	}

	for _, v := range day2Costs {
		if fc2 == 0 {
			fc2 = v
		}
		if v < fc2 {
			fc2 = v
		}

	}

	fmt.Printf("Part 1: Fuel Cost %v\n", fc1)
	fmt.Printf("Part 2: Fuel Cost %v\n", fc2)

}
