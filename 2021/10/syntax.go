package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var opens = map[string]string{"[": "]", "(": ")", "{": "}", "<": ">"}
var closes = map[string]int{"]": 57, ")": 3, "}": 1197, ">": 25137}

type Line []string

func (l Line) String() string {
	f := ""
	for _, v := range l {
		f += fmt.Sprintf("%s", v)
	}
	return f
}

func getText(f string) [][]string {
	allStrings := [][]string{}
	content, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	// trim trailing newline
	line := strings.TrimSuffix(string(content), "\n")

	for _, v := range strings.Split(line, "\n") {
		var row []string

		for _, char := range strings.Split(v, "") {
			row = append(row, char)
		}
		allStrings = append(allStrings, row)
	}
	return allStrings

}

func findFirstProblem(t Line) (string, int) {
	counter := []string{}
	for i, v := range t {
		nextClose, isOpen := opens[v]
		if isOpen == true {
			counter = append(counter, nextClose)
			continue
		}

		thisClose, isClose := closes[v]
		if isClose == true {
			// if the closing character is the top on the stack
			if v == counter[len(counter)-1] {
				counter = counter[:len(counter)-1]
				continue
			} else {
				fmt.Printf("\tfound point value %d at %d: WRONG. Should be %s\n", thisClose, i, counter[len(counter)-1])
				return v, thisClose
			}
		}
	}
	return "", 0
}

func main() {
	allChars := getText("input.txt")
	allPoints := 0
	for _, row := range allChars {
		_, points := findFirstProblem(row)
		allPoints += points

	}
	fmt.Println(allPoints)

}
