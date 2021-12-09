package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getCrabs(f string) []int {
	content, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	line := strings.TrimSuffix(string(content), "\n")

	// initialize an empty map
	var allCrabs []int

	// populate the slice with the crabs
	for _, v := range strings.Split(line, ",") {
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

func main() {
	allCrabs := getCrabs("crabs.txt")
	sort.Ints(allCrabs)
	fuelCosts := map[int]int{}

	// Fuel costs 1 per space
	for j := GetMin(&allCrabs); j <= GetMax(&allCrabs); j++ {
		for _, v := range allCrabs {
			// fuelCosts[j] += Abs(v, j)
			fuelCosts[j] += CalcFuelCost(Abs(v, j))
		}
	}

	var fc int
	// find
	for k, v := range fuelCosts {
		if fc == 0 {
			fc = v
		}
		if v < fc {
			fc = v
			fmt.Printf("Position %v: Fuel Cost %v\n", k, v)
		}
	}

}
