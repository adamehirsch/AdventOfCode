package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type depthMap struct {
	dm         [][]int
	riskLevels []int
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
	fmt.Print(getMap("depths.txt"))
}
