package utils

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/fatih/color"
)

// var red = color.New(color.FgRed).SprintFunc()

var green = color.New(color.FgGreen, color.Bold).PrintfFunc()

type Point struct {
	X int
	Y int
}

// func (p Point) String() string {
// 	return fmt.Sprintf("{%d %d}", p.X, p.Y)
// }

type GridMap [][]int

func (gm GridMap) String() string {
	sb := ""
	for y, row := range gm {
		for x := 0; x < len(row); x++ {
			if gm[y][x] == math.MaxInt {
				sb += color.GreenString(fmt.Sprintf("%5s", "+"))
			} else {
				sb += fmt.Sprintf("%5d", gm[y][x])
			}
		}
		sb += "\n"
	}
	return sb
}

func (gm *GridMap) PrintWinningPath(wp []Point) {
	for y, row := range *gm {
		for x := 0; x < len(row); x++ {
			if ContainsPoint(wp, Point{X: x, Y: y}) {
				green("%2d*", (*gm)[y][x])
			} else {
				fmt.Printf("%3d", (*gm)[y][x])
			}
		}
		fmt.Print("\n")
	}
}

func (gm *GridMap) PointValue(p Point) int {
	return (*gm)[p.Y][p.X]
}

func (gm *GridMap) SetValue(p Point, v int) {
	(*gm)[p.Y][p.X] = v
}

func MakeGridMap(Xmax, Ymax, def int) GridMap {
	gm := make([][]int, Ymax+1)
	for y := 0; y <= Ymax; y++ {
		gm[y] = make([]int, Xmax+1)
		for x := range gm[y] {
			gm[y][x] = def
		}
	}
	return gm
}

func Neighbors(b GridMap, x, y int, eight bool) []Point {

	// zero value for bool is false
	var ym, xm, yp, xp bool

	// catch the case where someone asks for a point on the border of the game
	if y > 0 {
		ym = true
	}
	if x > 0 {
		xm = true
	}
	if y < len(b)-1 {
		yp = true
	}
	if x < len(b)-1 {
		xp = true
	}

	neighbors := []Point{}

	if xm && ym && eight {
		neighbors = append(neighbors, Point{X: x - 1, Y: y - 1})
	}
	if ym {
		neighbors = append(neighbors, Point{X: x, Y: y - 1})
	}
	if xp && ym && eight {
		neighbors = append(neighbors, Point{X: x + 1, Y: y - 1})
	}
	if xp {
		neighbors = append(neighbors, Point{X: x + 1, Y: y})
	}
	if xp && yp && eight {
		neighbors = append(neighbors, Point{X: x + 1, Y: y + 1})
	}
	if yp {
		neighbors = append(neighbors, Point{X: x, Y: y + 1})
	}
	if xm && yp && eight {
		neighbors = append(neighbors, Point{X: x - 1, Y: y + 1})
	}
	if xm {
		neighbors = append(neighbors, Point{X: x - 1, Y: y})
	}

	return neighbors

}

func EightNeighbors(b GridMap, x, y int) []Point {
	return Neighbors(b, x, y, true)
}

func FourNeighbors(b GridMap, x, y int) []Point {
	return Neighbors(b, x, y, false)
}

type BoolMap [][]bool

func (bm BoolMap) String() string {
	sb := ""
	for y, row := range bm {
		for x := 0; x < len(row); x++ {
			if bm[y][x] {
				sb += color.GreenString("+")
			} else {
				sb += color.RedString("-")
			}
		}
		sb += "\n"
	}
	return sb
}

func MakeBoolMap(Xmax, Ymax int) BoolMap {
	gm := make([][]bool, Ymax+1)
	for y := 0; y <= Ymax; y++ {
		gm[y] = make([]bool, Xmax+1)
		for x := range gm[y] {
			gm[y][x] = false
		}
	}
	return gm
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func Reverse(s []string) []string {
	r := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}

func Contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func ContainsPoint(slice []Point, val Point) bool {
	for _, p := range slice {
		if p.X == val.X && p.Y == val.Y {
			return true
		}
	}
	return false
}

func Opener(f string, isData bool) (string, error) {

	if isData {
		pwd, _ := os.Getwd()
		f = filepath.Join(pwd, f)
	}

	file, err := os.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.TrimSuffix(string(file), "\n"), nil
}

func GetGridMap(f string) GridMap {
	var dm [][]int

	content, err := Opener(f, true)
	if err != nil {
		log.Fatal(err)
	}
	// trim trailing newline

	for _, v := range strings.Split(content, "\n") {
		var row []int

		for _, char := range strings.Split(v, "") {
			f, _ := strconv.Atoi(char)
			row = append(row, f)
		}
		dm = append(dm, row)
	}
	return dm

}
