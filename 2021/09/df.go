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

func main() {
	dm := getMap("depths.txt")
	for y := 0; y < len(dm.dm); y++ {
		for x := 0; x < len(dm.dm[0]); x++ {
			dm.riskLevels += dm.GetRiskLevel(x, y)
		}
	}
	fmt.Printf("Risk Levels Total: %d\n", dm.riskLevels)

}
