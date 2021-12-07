package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type grid [][]int

type point struct {
	X int
	Y int
}

func (p point) String() string {
	return fmt.Sprintf("(X: %d, Y: %d)", p.X, p.Y)
}

type line struct {
	begin point
	end   point
}

func (l line) String() string {
	return fmt.Sprintf("%v -> %v", l.begin, l.end)
}

func getCoords(input string) []line {
	file, err := os.Open(input)

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	var allLines []line

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var newLine line

		b := strings.Split(scanner.Text(), " -> ")
		fmt.Println(b)

		if len(b) != 2 {
			fmt.Println("THIS SHOULD NOT HAPPEN")
			continue
		}

		for i, v := range b {
			// first point, then second point
			c := strings.Split(v, ",")

			d, _ := strconv.Atoi(c[0])
			e, _ := strconv.Atoi(c[1])
			if i == 0 {
				newLine.begin.X = d
				newLine.begin.Y = e
			} else if i == 1 {
				newLine.end.X = d
				newLine.end.Y = e
				fmt.Println("\t", newLine.end)
			}

			if i == 1 {
				allLines = append(allLines, newLine)
			}
		}

	}
	return allLines
}

func main() {
	allLines := getCoords("input.txt")
	for _, v := range allLines {
		fmt.Println(v)
	}
}
