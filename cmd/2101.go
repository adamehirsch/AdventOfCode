package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2101Cmd = &cobra.Command{
	Use:                   "day2101",
	Short:                 "2021 Advent of Code Day 1",
	DisableFlagsInUseLine: true,
	Run:                   day2101Func,
}

func init() {
	rootCmd.AddCommand(day2101Cmd)
}

func day2101Func(cmd *cobra.Command, args []string) {

	file, err := utils.Opener("data/2101.txt", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var a, b, totalDeeper, i int

	for _, v := range strings.Split(file, "\n") {
		b, _ = strconv.Atoi(v)

		// skip the first line, total all the deeper spots
		if i > 0 && b > a {
			totalDeeper++
		}
		a = b
		i++
	}

	fmt.Printf("Total Deeper: %d\n", totalDeeper)
}
