package cmd

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2109Cmd = &cobra.Command{
	Use:                   "day2109",
	Short:                 "2021 Advent of Code Day 9",
	DisableFlagsInUseLine: true,
	Run:                   day2109Func,
}

func init() {
	rootCmd.AddCommand(day2109Cmd)
}

type depthMap struct {
	dm         [][]int
	riskLevels int
}

func (dm *depthMap) GetNeighbors(x, y int) ([]int, error) {
	results := []int{}

	if y > len(dm.dm) || x > len(dm.dm[0]) || x < 0 || y < 0 {
		return []int{}, errors.New("invalid coordinate specified")
	}

	if y > 0 {
		results = append(results, dm.dm[y-1][x])
	}
	if x != len(dm.dm[0])-1 {
		results = append(results, dm.dm[y][x+1])
	}
	if y != len(dm.dm)-1 {
		results = append(results, dm.dm[y+1][x])
	}
	if x > 0 {
		results = append(results, dm.dm[y][x-1])
	}

	return results, nil
}

func (dm *depthMap) GetSizes() (int, int) {
	return len(dm.dm) - 1, len(dm.dm[0]) - 1
}

func (dm *depthMap) GetRiskLevel(x, y int) int {
	neighbors, _ := dm.GetNeighbors(x, y)

	center := dm.dm[y][x]

	for _, v := range neighbors {
		if center >= v {
			return 0
		}
	}
	return center + 1
}

func getBasinMap(f string) depthMap {
	var dm depthMap
	file, err := utils.Opener("data/2109.txt", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, v := range strings.Split(file, "\n") {
		var row []int

		for _, char := range strings.Split(v, "") {
			f, _ := strconv.Atoi(char)
			row = append(row, f)
		}
		dm.dm = append(dm.dm, row)
	}
	return dm

}

func alreadyScanned(x, y int, ts *[]Point) bool {
	// skip any already scanned points
	for _, v := range *ts {
		if v.X == x && v.Y == y {
			return true
		}
	}
	return false
}

func (dm *depthMap) FloodFill(x, y int, tb *int, s *[]Point) (int, error) {

	if y > len(dm.dm) || x > len(dm.dm[0]) || x < 0 || y < 0 {
		return -1, errors.New("invalid coordinate specified")
	}

	if alreadyScanned(x, y, s) {
		return 0, nil
	}

	// record that we've scanned this point already
	*s = append(*s, Point{X: x, Y: y})

	// the basins are bounded by 9-values
	if dm.dm[y][x] == 9 {
		return 0, nil
	} else {
		// the square we're on is in a basin
		*tb += 1
	}

	if y > 0 {
		dm.FloodFill(x, y-1, tb, s)
	}
	if x != len(dm.dm[0])-1 {
		dm.FloodFill(x+1, y, tb, s)
	}
	if y != len(dm.dm)-1 {
		dm.FloodFill(x, y+1, tb, s)
	}
	if x > 0 {
		dm.FloodFill(x-1, y, tb, s)
	}

	return *tb, nil
}

func day2109Func(cmd *cobra.Command, args []string) {

	dm := getBasinMap("depths.txt")
	for y := 0; y < len(dm.dm); y++ {
		for x := 0; x < len(dm.dm[0]); x++ {
			dm.riskLevels += dm.GetRiskLevel(x, y)
		}
	}
	fmt.Printf("Part 1: Risk Levels Total: %d\n", dm.riskLevels)

	toScan := []Point{}
	basins := []int{}

	for y := 0; y < len(dm.dm); y++ {
		for x := 0; x < len(dm.dm[0]); x++ {
			if !alreadyScanned(x, y, &toScan) {
				tb := 0
				bv, _ := dm.FloodFill(x, y, &tb, &toScan)
				if bv != 0 {
					basins = append(basins, bv)
				}

			}

		}
	}

	sort.Ints(basins)
	biggestIndex := len(basins) - 1
	fmt.Printf("Part 2 solution: %d\n", basins[biggestIndex]*basins[biggestIndex-1]*basins[biggestIndex-2])

}
