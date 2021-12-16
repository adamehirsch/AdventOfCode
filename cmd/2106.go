package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2106Cmd = &cobra.Command{
	Use:                   "day2106",
	Short:                 "2021 Advent of Code Day 6",
	DisableFlagsInUseLine: true,
	Run:                   day2106Func,
}

func init() {
	rootCmd.AddCommand(day2106Cmd)
}

type fishingGround struct {
	fg       map[int]int64
	daycount int
}

func (fg *fishingGround) Sunrise() {
	zeroes := fg.fg[0]

	fg.fg[0] = fg.fg[1]
	fg.fg[1] = fg.fg[2]
	fg.fg[2] = fg.fg[3]
	fg.fg[3] = fg.fg[4]
	fg.fg[4] = fg.fg[5]
	fg.fg[5] = fg.fg[6]
	fg.fg[6] = zeroes + fg.fg[7]

	fg.fg[7] = fg.fg[8]
	fg.fg[8] = zeroes

	fg.daycount++

}

func (fg *fishingGround) Total() int64 {
	a := int64(0)
	a += fg.fg[0]
	a += fg.fg[1]
	a += fg.fg[2]
	a += fg.fg[3]
	a += fg.fg[4]
	a += fg.fg[5]
	a += fg.fg[6]
	a += fg.fg[7]
	a += fg.fg[8]

	return a
}

func (fg fishingGround) String() string {
	return fmt.Sprintf("0: %d\n1: %d\n2: %d\n3: %d\n4: %d\n5: %d\n6: %d\n7: %d\n8: %d\n\ndc: %d\n\n", fg.fg[0], fg.fg[1], fg.fg[2], fg.fg[3], fg.fg[4], fg.fg[5], fg.fg[6], fg.fg[7], fg.fg[8], fg.daycount)

}

func getFish(f string) map[int]int64 {
	file, err := utils.Opener(f, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// initialize an empty map
	allFish := map[int]int64{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}

	// populate the map with the total numbers of fish, keyed by counter value
	for _, v := range strings.Split(file, ",") {
		f, _ := strconv.Atoi(v)
		allFish[f] += 1
	}
	return allFish
}

func day2106Func(cmd *cobra.Command, args []string) {
	allFish := getFish("data/2106.txt")
	fg := fishingGround{fg: allFish, daycount: 0}

	for i := 0; i < 256; i++ {
		fg.Sunrise()
		if i == 80-1 {
			fmt.Println("Part 1: 80 cycles ", fg.Total())
		}
	}
	fmt.Println("Part 2: 256 cycles ", fg.Total())

}
