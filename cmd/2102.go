package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2102Cmd = &cobra.Command{
	Use:                   "day2102",
	Short:                 "2021 Advent of Code Day 2",
	DisableFlagsInUseLine: true,
	Run:                   day2102Func,
}

func init() {
	rootCmd.AddCommand(day2102Cmd)
}

func day2102Func(cmd *cobra.Command, args []string) {

	file, err := utils.Opener("data/2102.txt", true)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var r, d, d1, a int // range, depth, aim, index

	for _, v := range strings.Split(file, "\n") {

		s := strings.Fields(v)

		direction := s[0]
		distance, _ := strconv.Atoi(s[1])

		switch direction {
		case "forward":
			r += distance
			d += a * distance
		case "up":
			a -= distance
			d1 -= distance
		case "down":
			a += distance
			d1 += distance
		default:
			fmt.Println("THIS SHOULD NOT HAPPEN")
		}

	}
	fmt.Printf("Part 1: Range * Depth is %d\n", r*d1)

	fmt.Printf("Part 2: Range * Depth is %d\n", r*d)

}
