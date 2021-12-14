package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("depths.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var a, b, totalDeeper, i int

	for scanner.Scan() {
		b, _ = strconv.Atoi(scanner.Text())

		// skip the first line
		if i > 0 && b > a {
			totalDeeper++
		}
		fmt.Printf("%d - %d: %d %t\n", a, b, totalDeeper, b > a)

		a = b
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
