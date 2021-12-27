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

func Dijkstra(m, costMap *utils.GridMap, bestPath *[]utils.Point, finish utils.Point, seen *[]utils.Point, toCheck *[]utils.Point) {
	// pull the working point off the toCheck list
	workingPoint := (*toCheck)[0]

	startCost := (*costMap).PointValue(workingPoint)

	// Mark this origin point Seen
	if !utils.ContainsPoint(*seen, workingPoint) {
		*seen = append(*seen, workingPoint)
	}

	// take the working point off the queue
	*toCheck = (*toCheck)[1:]

	for _, p := range utils.FourNeighbors(*m, workingPoint.X, workingPoint.Y) {

		if !utils.ContainsPoint(*seen, p) {

			if !utils.ContainsPoint(*toCheck, p) {
				// for any of the neighbors we haven't seen, let's check them
				*toCheck = append(*toCheck, p)
			}

			// for each of the neighbors, set its cost to the cost of (this origin point + the node's current cost)
			if (*costMap).PointValue(p) > startCost+(*m).PointValue(p) {
				(*costMap).SetValue(p, startCost+(*m).PointValue(p))
			}

		}
	}
	// sort the neighbors by their current cost
	sort.Slice(*toCheck, func(i, j int) bool {
		return (*costMap)[(*toCheck)[i].Y][(*toCheck)[i].X] < (*costMap)[(*toCheck)[j].Y][(*toCheck)[j].X]
	})

}

func day2115Func(cmd *cobra.Command, args []string) {

	caveMap := utils.GetGridMap("data/2115-sample.txt")
	costMap := utils.MakeGridMap(len(caveMap[0])-1, len(caveMap)-1, math.MaxInt)

	seenPoints := []utils.Point{}
	toCheck := []utils.Point{{X: 0, Y: 0}}

	bestPath := []utils.Point{}

	// the cost of the starting point is 0
	costMap[0][0] = 0

	for len(toCheck) > 0 {
		Dijkstra(&caveMap, &costMap, &bestPath, utils.Point{X: len(caveMap[0]) - 1, Y: len(caveMap) - 1}, &seenPoints, &toCheck)

	}

	fmt.Print(caveMap)
	fmt.Println(toCheck)
	fmt.Print(costMap)
}
