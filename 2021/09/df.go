package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type depthMap struct {
	dm         [][]int
	riskLevels int
}

type Point struct {
	X int
	Y int
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

func getMap(f string) depthMap {
	var dm depthMap
	content, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	// trim trailing newline
	line := strings.TrimSuffix(string(content), "\n")

	for _, v := range strings.Split(line, "\n") {
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
		fmt.Printf("NINE: %d, %d: %d (tb: %d)\n", x, y, dm.dm[y][x], tb)
		return 0, nil
	} else {
		*tb += 1
		// the square we're on is in a basin
		fmt.Printf("HIT: %d, %d: %d (tb: %d)\n", x, y, dm.dm[y][x], tb)
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
		dm.FloodFill(x-1, y-1, tb, s)
	}

	return *tb, nil
}

func main() {
	dm := getMap("sampledepths.txt")
	for y := 0; y < len(dm.dm); y++ {
		for x := 0; x < len(dm.dm[0]); x++ {
			dm.riskLevels += dm.GetRiskLevel(x, y)
		}
	}
	fmt.Printf("Risk Levels Total: %d\n", dm.riskLevels)

	toScan := []Point{}

	for y := 0; y < len(dm.dm); y++ {
		for x := 0; x < len(dm.dm[0]); x++ {
			if alreadyScanned(x, y, &toScan) == false {
				fmt.Printf("STARTING NEW BASIN: %d, %d\n", x, y)
				tb := 0
				bv, _ := dm.FloodFill(x, y, &tb, &toScan)
				if bv != 0 {
					fmt.Printf("%d, %d: %d\n", x, y, bv)

				}
			}

		}
	}
	fmt.Print(toScan)

}
