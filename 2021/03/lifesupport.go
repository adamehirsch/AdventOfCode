package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	if most[0] == true {
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

func main() {

	const lineLength = 12

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var allLines []string

	for scanner.Scan() {
		allLines = append(allLines, scanner.Text())
	}

	oxygenLines := allLines
	co2lines := allLines

	var prefix string

	// Oxygen generator rating
	for num := 0; num < lineLength; num++ {
		prefix += findMostLeast(oxygenLines, num, "1")
		oxygenLines = selectLines(oxygenLines, prefix)
		fmt.Printf("O2  all: %d, remaining: %d\n", len(allLines), len(oxygenLines))
		if len(oxygenLines) == 1 {
			break
		}
	}

	// CO2 generator rating
	prefix = ""
	for num := 0; num < lineLength; num++ {
		prefix += findMostLeast(co2lines, num, "0", false)
		co2lines = selectLines(co2lines, prefix)
		fmt.Printf("CO2 all: %d, remaining: %d\n", len(allLines), len(co2lines))
		if len(co2lines) == 1 {
			break
		}
	}
	oxy, _ := strconv.ParseInt(oxygenLines[0], 2, 64)
	co2, _ := strconv.ParseInt(co2lines[0], 2, 64)

	fmt.Printf("%d * %d = %d", oxy, co2, oxy*co2)
	//

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
