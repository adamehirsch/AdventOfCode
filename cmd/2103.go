package cmd

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2103Cmd = &cobra.Command{
	Use:                   "day2103",
	Short:                 "2021 Advent of Code Day 3",
	DisableFlagsInUseLine: true,
	Run:                   day2103Func,
}

func init() {
	rootCmd.AddCommand(day2103Cmd)
}

func selectLines(s []string, p string) []string {
	c := make([]string, 0)
	for _, v := range s {
		if strings.HasPrefix(v, p) {
			c = append(c, v)
		}
	}
	return c
}

func findMostLeast(s []string, p int, tb string, most ...bool) string {
	// an awkward way to make an optional argument with default-true
	if len(most) == 0 {
		most = append(most, true)
	}

	fc := make(map[string]int)
	for _, v := range s {
		// keep a count of every character at index p
		fc[string(v[p])]++
	}

	// if all the values have the same count, return the tiebreaker
	if hasDupes(fc) {
		return tb
	}

	// otherwise return the string that had the highest count
	var rv string
	c := 0
	// look at the characters in the frequency count, return the key with the greatest count
	for k, v := range fc {
		if v > c {
			c = v
			rv = k
		}
	}

	if most[0] {
		return rv
	}
	// if we're looking for the least-frequent. This could be DRYer, I think
	for k, v := range fc {
		if v < c {
			c = v
			rv = k
		}
	}
	return rv
}

func hasDupes(m map[string]int) bool {
	// quick hack to see whether two values in a map are identical; in this case in the frequency counter
	x := make(map[int]struct{})
	for _, v := range m {
		if _, has := x[v]; has {
			return true
		}
		x[v] = struct{}{}
	}
	return false
}

func day2103Func(cmd *cobra.Command, args []string) {

	const lineLength = 12

	file, err := utils.Opener("data/2103.txt", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	allLines := strings.Split(file, "\n")

	oxygenLines := allLines
	co2lines := allLines

	var i int // index
	var gamma int
	var epsilon int

	fc := make(map[int]int) // frequency count accumulator

	// Part 1
	for _, v := range allLines {

		s := strings.Split(v, "")

		// for each index position on each line, count the 1
		for m := 0; m < len(s); m++ {
			n, _ := strconv.Atoi(s[m])
			if n == 1 {
				fc[m]++
			}
		}
		i++
	}
	// walk the frequency counter
	for c := 0; c < len(fc); c++ {

		// if more than half of the digits at that index were 1, add 2 to the power of MATH
		// ... the MATH part involved an off by one error around the length of the frequency counter

		a := int(math.Pow(2, float64(len(fc)-c-1)))

		if fc[c] > i/2 {
			gamma += a
		} else {
			epsilon += a
		}
	}
	var prefix string
	fmt.Printf("Part 1: %b * %b = %d\n", gamma, epsilon, gamma*epsilon)

	// Oxygen generator rating
	for num := 0; num < lineLength; num++ {
		prefix += findMostLeast(oxygenLines, num, "1")
		oxygenLines = selectLines(oxygenLines, prefix)
		// fmt.Printf("O2  all: %d, remaining: %d\n", len(allLines), len(oxygenLines))
		if len(oxygenLines) == 1 {
			break
		}
	}

	// CO2 generator rating
	prefix = ""
	for num := 0; num < lineLength; num++ {
		prefix += findMostLeast(co2lines, num, "0", false)
		co2lines = selectLines(co2lines, prefix)
		// fmt.Printf("CO2 all: %d, remaining: %d\n", len(allLines), len(co2lines))
		if len(co2lines) == 1 {
			break
		}
	}
	oxy, _ := strconv.ParseInt(oxygenLines[0], 2, 64)
	co2, _ := strconv.ParseInt(co2lines[0], 2, 64)

	fmt.Printf("Part 2: %d * %d = %d\n", oxy, co2, oxy*co2)
	//

}
