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

type Recorder struct {
	pairs   map[string]int64
	letters map[string]int64
}

func (r *Recorder) Step(pi ProteinInstructions) {
	newPairCounts := make(map[string]int64)
	// look at each of the protein-instructions
	for pair, insertion := range pi {
		if startingPairCount, found := r.pairs[pair]; found {
			// we've got a pair that needs inserting, iykwim
			newPairs := []string{string(pair[0]) + insertion, insertion + string(pair[1])}
			for _, np := range newPairs {
				// each of the New Pairs increments by the number of the original paircount
				newPairCounts[np] += startingPairCount
			}
			// each of the initial letters gets incremented by the number of insertions
			r.letters[insertion] += startingPairCount
		}

	}
	r.pairs = newPairCounts
}

func StringPairs(s string) []string {
	// using this just once for initializing
	pairs := []string{}
	for i := 0; i < len(s); i++ {
		if i == 0 {
			continue
		}
		pairs = append(pairs, s[i-1:i+1])
	}
	return pairs
}

func getPairInstructions(f string) (Recorder, ProteinInstructions) {
	file, err := utils.Opener(f, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pi := ProteinInstructions{}
	r := Recorder{pairs: make(map[string]int64), letters: make(map[string]int64)}

	for _, v := range strings.Split(file, "\n") {
		s := strings.Split(v, " -> ")
		if len(s) == 2 {
			// these are the insertion instructions
			pi[s[0]] = s[1]

		} else if len(v) > 0 {
			// these are the initial pairs
			for _, p := range StringPairs(v) {
				r.pairs[p]++
			}

			for _, t := range v {
				r.letters[string(t)]++
			}

		}

	}
	return r, pi

}

func FindSolution(r Recorder) int64 {
	var max, min int64
	for _, v := range r.letters {

		if min == 0 {
			min = v
		}

		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}
	return max - min
}

func day2114Func(cmd *cobra.Command, args []string) {
	recorder, instructions := getPairInstructions("data/2114.txt")

	for i := 0; i < 40; i++ {
		recorder.Step(instructions)
		if i == 9 || i == 39 {
			fmt.Printf("Step %d: %d\n", i+1, FindSolution(recorder))
		}

	}
}
