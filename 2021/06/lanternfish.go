package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

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
	content, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	line := strings.TrimSuffix(string(content), "\n")

	// initialize an empty map
	allFish := map[int]int64{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}

	// populate the map with the total numbers of fish, keyed by counter value
	for _, v := range strings.Split(line, ",") {
		f, _ := strconv.Atoi(v)
		allFish[f] += 1
	}
	return allFish
}

func main() {
	allFish := getFish("initial-fish.txt")
	fg := fishingGround{fg: allFish, daycount: 0}

	for i := 0; i < 256; i++ {
		roundStart := time.Now()
		fg.Sunrise()
		if math.Mod(float64(i), float64(20)) == 0 {
			fmt.Println(fg)
			fmt.Println(time.Since(roundStart))
		}
	}
	fmt.Print(fg.Total())

}
