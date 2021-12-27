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

func Dijkstra(m, costMap *utils.GridMap, bestMap *PathMap, finish utils.Point, seen *[]utils.Point, toCheck *[]utils.Point) {
	// pull the working point off the toCheck list
	workingPoint := (*toCheck)[0]
	// // the current best path to the working point
	// wp_path := (*bestMap)[workingPoint.Y][workingPoint.X]

	// // if we've
	// if wp_path[len(wp_path)-1].X != workingPoint.X || wp_path[len(wp_path)-1].Y != workingPoint.Y {
	// 	fmt.Println("1: ADDING ", wp_path, workingPoint)
	// 	(*bestMap)[workingPoint.Y][workingPoint.X] = append(wp_path, workingPoint)
	// }

	startCost := (*costMap).PointValue(workingPoint)
	startPath := (*bestMap)[workingPoint.Y][workingPoint.X]

	// fmt.Printf("WP: %v BP: %v\n", workingPoint, startPath)
	// fmt.Printf("1,2 bestPath: %v\n", (*bestMap)[2][1])
	// fmt.Printf("0,3 bestPath: %v\n", (*bestMap)[3][0])

	// Mark this origin point Seen
	if !utils.ContainsPoint(*seen, workingPoint) {
		*seen = append(*seen, workingPoint)
	}

	// take the working point off the queue
	*toCheck = (*toCheck)[1:]

	for _, neighbor := range utils.FourNeighbors(*m, workingPoint.X, workingPoint.Y) {

		// fmt.Printf("  NEIGHBORS: wp %v -> %v\n", workingPoint, neighbor)

		if !utils.ContainsPoint(*seen, neighbor) {

			// fmt.Printf("  !SEEN: wp %v -> %v\n", workingPoint, neighbor)

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
				//  = newPath

				// fmt.Printf("     AFTER %v %v\n", neighbor, (*bestMap)[neighbor.Y][neighbor.X])
				// fmt.Printf("     1,2 bestPath: %v\n", (*bestMap)[2][1])
				// fmt.Printf("     0,3 bestPath: %v\n", (*bestMap)[3][0])
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

	seenPoints := []utils.Point{}
	toCheck := []utils.Point{{X: 0, Y: 0}}

	// the cost of the starting point is 0
	costMap[0][0] = 0
	bestMap := MakePathMap(Xmax, Ymax)
	bestMap[0][0] = append(bestMap[0][0], utils.Point{X: 0, Y: 0})

	for len(toCheck) > 0 {
		Dijkstra(&caveMap, &costMap, &bestMap, utils.Point{X: len(caveMap[0]) - 1, Y: len(caveMap) - 1}, &seenPoints, &toCheck)

	}

	caveMap.PrintWinningPath(bestMap[Ymax][Xmax])

	fmt.Printf("Part 1: Total danger at destination: %d\n", costMap[Ymax][Xmax])
}
