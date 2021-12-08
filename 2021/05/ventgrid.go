package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type grid [][]int

func newGrid(size int) grid {
	g := make(grid, size)
	for i := 0; i < size; i++ {
		g[i] = make([]int, size)
	}
	return g
}

func (g *grid) countWinners() int {
	winners := 0
	for _, v := range *g {
		for _, val := range v {
			if val > 1 {
				winners += 1
			}
		}
	}
	return winners
}

func (g *grid) drawLine(l line) {

	if l.isOrthogonal() {
		// is there a more elegant way to get these in ascending order? almost certainly
		// do I care? probably not.

		bx, ex := l.begin.X, l.end.X
		if bx > ex {
			ex, bx = l.begin.X, l.end.X
		}

		by, ey := l.begin.Y, l.end.Y
		if by > ey {
			ey, by = l.begin.Y, l.end.Y
		}

		for bx := bx; bx <= ex; bx++ {
			for by := by; by <= ey; by++ {
				(*g)[by][bx] += 1
			}
		}
	} else {

		bx, ex := l.begin.X, l.end.X
		by, ey := l.begin.Y, l.end.Y

		mx := 1
		my := 1

		if ex < bx {
			mx = -1
		}

		if ey < by {
			my = -1
		}

		// Since these lines are strictly diagonal, the range = run, so I can just
		// make as many steps as the x-grid
		for i := 0; i <= int(math.Abs(float64(bx-ex))); i++ {
			(*g)[by+(i*my)][bx+(i*mx)] += 1
		}

	}
}

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

func (l line) isOrthogonal() bool {
	if l.begin.X == l.end.X || l.begin.Y == l.end.Y {
		return true
	}
	return false
}

func getLines(input string) []line {
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
			}

			if i == 1 {
				allLines = append(allLines, newLine)
			}
		}

	}
	return allLines
}

func main() {
	allLines := getLines("input.txt")

	thisGrid := newGrid(1000)

	for _, l := range allLines {
		thisGrid.drawLine(l)

	}
	fmt.Println(thisGrid.countWinners())
}
