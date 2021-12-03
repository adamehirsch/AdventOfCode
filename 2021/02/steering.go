package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("dirs.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var r, d, i int // range, depth, index

	for scanner.Scan() {

		s := strings.Fields(scanner.Text())

		u, _ := strconv.Atoi(s[1])

		direction, distance := s[0], u

		switch direction {
		case "forward":
			r += distance
		case "up":
			d -= distance
		case "down":
			d += distance
		default:
			fmt.Println("THIS SHOULD NOT HAPPEN")
		}

		i++
	}

	fmt.Printf("Range * Depth is %d", r*d)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
