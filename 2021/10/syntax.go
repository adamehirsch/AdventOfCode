package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var opens = map[string]string{"[": "]", "(": ")", "{": "}", "<": ">"}
var closes = map[string]string{"]": "]", ")": ")", "{": "{", ">": ">"}

type Line []string

func (l Line) String() string {
	f := ""
	for _, v := range l {
		f += fmt.Sprintf("%s ", v)
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

func findFirstProblem(t Line) (string, bool) {
	counter := []string{}
	fmt.Println(t)
	for i, v := range t {

		fmt.Printf("%d: %s\n", i, v)
		nextClose, isOpen := opens[v]
		if isOpen == true {
			counter = append(counter, nextClose)
			fmt.Printf("\tAdding %s\n", counter)
			continue
		}

		thisClose, isClose := closes[v]
		if isClose == true {
			fmt.Printf("\tfound %s at %d: removing %s\n", thisClose, i, counter[len(counter)-1])
			// if the closing character is the top on the stack
			if thisClose == counter[len(counter)-1] {
				counter = counter[:len(counter)-1]
				continue
			} else {
				fmt.Printf("\tfound %s at %d: WRONG %s\n", thisClose, i, counter[len(counter)-1])

				return v, true
			}
		}
	}
	return "", false
}

func main() {
	allChars := getText("sample.txt")
	for _, row := range allChars {
		char, foundProblem := findFirstProblem(row)
		if foundProblem {
			for j := 0; j < len(row); j++ {
				fmt.Printf("%s ", row[j])
			}
			fmt.Print("\n")
			fmt.Println(char)
		}

	}

}
