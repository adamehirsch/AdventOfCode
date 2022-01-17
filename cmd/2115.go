package cmd

import (
	"fmt"
	"math"
	"sort"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2115Cmd = &cobra.Command{
	Use:                   "day2115",
	Short:                 "2021 Advent of Code Day 15",
	DisableFlagsInUseLine: true,
	Run:                   day2115Func,
}

func init() {
	rootCmd.AddCommand(day2115Cmd)
}

type PathMap [][][]utils.Point

func MakePathMap(xMax, yMax int) PathMap {
	pm := make([][][]utils.Point, yMax+1)
	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			pm[y] = make([][]utils.Point, xMax+1)
			pm[y][x] = make([]utils.Point, xMax+1)
		}
	}
	return pm
}

func FiveMap(origMap utils.GridMap) (utils.GridMap, utils.GridMap, PathMap) {

	origXmax := len(origMap[0]) - 1
	origYmax := len(origMap) - 1

	Xmax := (5 * len(origMap[0])) - 1
	Ymax := (5 * (len(origMap))) - 1

	newMap := utils.MakeGridMap(Xmax, Ymax, math.MaxInt)

	for y := 0; y <= Ymax; y++ {
		for x := 0; x <= Xmax; x++ {
			// figure out the coordinates in the "original" sector
			iy := math.Mod(float64(y), float64(origYmax+1))
			ix := math.Mod(float64(x), float64(origXmax+1))

			baseValue := origMap[int(iy)][int(ix)]

			// we are turning Golang's insistence on flooring any int/int to our advantage
			lateralDistance := x / (origXmax + 1)
			verticalDistance := y / (origYmax + 1)

			// set the value with the modified value from the base map
			newMap[y][x] = NumWrap(baseValue, lateralDistance+verticalDistance)

		}
	}

	costMap := utils.MakeGridMap(Xmax, Ymax, math.MaxInt)
	// the cost of the starting point is 0
	costMap[0][0] = 0
	bestMap := MakePathMap(Xmax, Ymax)
	bestMap[0][0] = append(bestMap[0][0], utils.Point{X: 0, Y: 0})
	return newMap, costMap, bestMap

}

func NumWrap(i, m int) int {
	z := i + m
	for z > 9 {
		z -= 9
	}
	return z
}

func Dijkstra(m, costMap *utils.GridMap, bestMap *PathMap, finish utils.Point, seen *map[utils.Point]bool, toCheck *[]utils.Point) {
	// pull the working point off the toCheck list
	workingPoint := (*toCheck)[0]

	startCost := (*costMap).PointValue(workingPoint)
	startPath := (*bestMap)[workingPoint.Y][workingPoint.X]

	// Mark this origin point Seen
	(*seen)[workingPoint] = true

	// take the working point off the queue
	*toCheck = (*toCheck)[1:]

	for _, neighbor := range utils.FourNeighbors(*m, workingPoint.X, workingPoint.Y) {

		// if !utils.ContainsPoint(*seen, neighbor) {

		if _, found := (*seen)[neighbor]; !found {
			if !utils.ContainsPoint(*toCheck, neighbor) {
				// for any of the neighbors we haven't seen, let's check them
				*toCheck = append(*toCheck, neighbor)
			}

			// for each of the neighbors, set its cost to the cost of (this origin point + the node's current cost)
			if (*costMap).PointValue(neighbor) > startCost+(*m).PointValue(neighbor) {

				(*costMap).SetValue(neighbor, startCost+(*m).PointValue(neighbor))

				// copy the map from the working point onto this destination point, appending itself as the best path
				nc := []utils.Point{}
				newPath := append(nc, startPath...)

				(*bestMap)[neighbor.Y][neighbor.X] = append(newPath, utils.Point{X: neighbor.X, Y: neighbor.Y})

			}

		}
	}
	// sort the neighbors by their current cost
	sort.Slice(*toCheck, func(i, j int) bool {
		return (*costMap)[(*toCheck)[i].Y][(*toCheck)[i].X] < (*costMap)[(*toCheck)[j].Y][(*toCheck)[j].X]
	})

}

func day2115Func(cmd *cobra.Command, args []string) {

	caveMap := utils.GetGridMap("data/2115.txt")
	Xmax := len(caveMap[0]) - 1
	Ymax := len(caveMap) - 1

	costMap := utils.MakeGridMap(Xmax, Ymax, math.MaxInt)

	// seenPoints := []utils.Point{}
	seenMap := make(map[utils.Point]bool)

	// toCheck := []utils.Point{{X: 0, Y: 0}}
	toCheck := make([]utils.Point, 0, Xmax*Ymax)
	toCheck = append(toCheck, utils.Point{X: 0, Y: 0})

	// the cost of the starting point is 0
	costMap[0][0] = 0
	bestMap := MakePathMap(Xmax, Ymax)
	bestMap[0][0] = append(bestMap[0][0], utils.Point{X: 0, Y: 0})

	for len(toCheck) > 0 {
		Dijkstra(&caveMap, &costMap, &bestMap, utils.Point{X: len(caveMap[0]) - 1, Y: len(caveMap) - 1}, &seenMap, &toCheck)

	}

	// caveMap.PrintWinningPath(bestMap[Ymax][Xmax])

	fmt.Printf("Part 1: Total danger at destination: %d\n", costMap[Ymax][Xmax])

	bigMap, bigCostMap, bigPathMap := FiveMap(caveMap)

	// bigPoints := []utils.Point{}
	bigPoints := make(map[utils.Point]bool)

	bigAgenda := make([]utils.Point, 0, Xmax*Ymax)
	bigAgenda = append(toCheck, utils.Point{X: 0, Y: 0})

	for len(bigAgenda) > 0 {
		Dijkstra(&bigMap, &bigCostMap, &bigPathMap, utils.Point{X: len(bigMap[0]) - 1, Y: len(bigMap) - 1}, &bigPoints, &bigAgenda)
	}

	// bigMap.PrintWinningPath(bigPathMap[len(bigMap[0])-1][len(bigMap)-1])

	fmt.Printf("Part 2: Total danger at destination: %d\n", bigCostMap[len(bigCostMap[0])-1][len(bigCostMap)-1])
}
