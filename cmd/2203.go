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

func init() {
	rootCmd.AddCommand(day2203Cmd)
}

func findCommonLetter(s string) string {
	l := strings.Split(s, "")
	firstHalf := l[:len(l)/2]
	secondHalf := l[len(l)/2:]
	fmt.Println(firstHalf)
	fmt.Println(secondHalf)
	fmt.Println("")
	return firstHalf[0]
}

func day2203Func(cmd *cobra.Command, args []string) {
	wholeInput, err := utils.Opener("data/2203-sample.txt", true)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := strings.Split(wholeInput, "\n")

	for _, line := range lines {
		findCommonLetter(line)
	}

}
