package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2114Cmd = &cobra.Command{
	Use:                   "day2114",
	Short:                 "2021 Advent of Code Day 14",
	DisableFlagsInUseLine: true,
	Run:                   day2114Func,
}

func init() {
	rootCmd.AddCommand(day2114Cmd)
}

type ProteinInstructions map[string]string

type Protein struct {
	p string
}

func getPairInstructions(f string) (Protein, ProteinInstructions) {
	file, err := utils.Opener(f, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pi := ProteinInstructions{}
	p := Protein{}
	for _, v := range strings.Split(file, "\n") {
		s := strings.Split(v, " -> ")
		if len(s) == 2 {
			pi[s[0]] = s[1]
		} else if len(v) > 0 {
			p.p = v
		}

	}
	return p, pi

}

func day2114Func(cmd *cobra.Command, args []string) {
	protein, pi := getPairInstructions("data/2114-sample.txt")
	fmt.Println(protein, pi)
}
