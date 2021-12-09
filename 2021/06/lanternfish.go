package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var bigOne = big.NewInt(1)

type fishingGround struct {
	fg       map[int]*big.Int
	daycount int
}

func (fg *fishingGround) Sunrise() {
	zeroes := big.NewInt(0)
	zeroes.Add(big.NewInt(0), fg.fg[0])

	fg.fg[0].Add(big.NewInt(0), fg.fg[1])
	fg.fg[1].Add(big.NewInt(0), fg.fg[2])
	fg.fg[2].Add(big.NewInt(0), fg.fg[3])
	fg.fg[3].Add(big.NewInt(0), fg.fg[4])
	fg.fg[4].Add(big.NewInt(0), fg.fg[5])
	fg.fg[5].Add(big.NewInt(0), fg.fg[6])
	fg.fg[6].Add(zeroes, fg.fg[7])
	fg.fg[7].Add(big.NewInt(0), fg.fg[8])

	fg.fg[8] = big.NewInt(0)

	for i := big.NewInt(0); i.Cmp(zeroes) < 0; i.Add(i, big.NewInt(1)) {
		fg.fg[8].Add(fg.fg[8], big.NewInt(1))
	}

	fg.daycount++

}

func (fg *fishingGround) Total() *big.Int {
	a := big.NewInt(0)
	a.Add(a, fg.fg[0])
	a.Add(a, fg.fg[1])
	a.Add(a, fg.fg[2])
	a.Add(a, fg.fg[3])
	a.Add(a, fg.fg[4])
	a.Add(a, fg.fg[5])
	a.Add(a, fg.fg[6])
	a.Add(a, fg.fg[7])
	a.Add(a, fg.fg[8])

	return a
}

func (fg fishingGround) String() string {
	return fmt.Sprintf("0: %s\n1: %s\n2: %s\n3: %s\n4: %s\n5: %s\n6: %s\n7: %s\n8: %s\n\ndc: %d\n\n", fg.fg[0].String(), fg.fg[1].String(), fg.fg[2].String(), fg.fg[3].String(), fg.fg[4].String(), fg.fg[5].String(), fg.fg[6].String(), fg.fg[7].String(), fg.fg[8].String(), fg.daycount)

}

func getFish(f string) map[int]*big.Int {
	content, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	line := strings.TrimSuffix(string(content), "\n")

	allFish := map[int]*big.Int{0: big.NewInt(0), 1: big.NewInt(0), 2: big.NewInt(0), 3: big.NewInt(0), 4: big.NewInt(0), 5: big.NewInt(0), 6: big.NewInt(0), 7: big.NewInt(0), 8: big.NewInt(0)}
	for _, v := range strings.Split(line, ",") {
		f, _ := strconv.Atoi(v)
		currValue := allFish[f]
		currValue = currValue.Add(currValue, bigOne)
		allFish[f] = currValue
	}
	return allFish
}

func main() {
	allFish := getFish("initial-fish.txt")
	fg := fishingGround{fg: allFish, daycount: 0}
	for i := 0; i < 80; i++ {
		fg.Sunrise()
	}
	fmt.Print(fg.Total().String())

}
