package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fatih/color"
)

func main() {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	file, err := os.Open("depths.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var a, b, c, d, r, i int

	for scanner.Scan() {
		d, _ = strconv.Atoi(scanner.Text())

		if i < 3 {
			slideWindow(&a, &b, &c, &d)
			i++
			continue
		}

		last_window := a + b + c
		this_window := b + c + d
		isDeeper := this_window > last_window

		if isDeeper == true {
			r++
		}

		deepResults := green(isDeeper)
		if isDeeper == true {
			deepResults = red(isDeeper)
		}

		fmt.Printf("%d: %d - %d: %d %s\n", i, last_window, this_window, r, deepResults)
		slideWindow(&a, &b, &c, &d)
		i++
	}

	fmt.Printf("%d deeper windows", r)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func slideWindow(a, b, c, d *int) {
	*a = *b
	*b = *c
	*c = *d
}
