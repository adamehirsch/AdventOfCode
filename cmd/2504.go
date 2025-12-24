package cmd

import (
	"fmt"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2504Cmd = &cobra.Command{
	Use:                   "day2504",
	Short:                 "2025 Advent of Code Day 04",
	DisableFlagsInUseLine: true,
	Run:                   day2504Func,
}

func init() {
	rootCmd.AddCommand(day2504Cmd)
}

func day2504Func(cmd *cobra.Command, args []string) {
	file, err := utils.Opener("data/2504.txt", true)
	if err != nil {
		fmt.Printf("failed to open input: %v\n", err)
		return
	}

	_ = file

	fmt.Println("TODO: implement day 2504")
}
