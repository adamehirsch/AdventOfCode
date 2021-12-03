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

	var r, d, a int // range, depth, aim

	for scanner.Scan() {

		s := strings.Fields(scanner.Text())

		direction := s[0]
		distance, _ := strconv.Atoi(s[1])

		switch direction {
		case "forward":
			r += distance
			d += a * distance
		case "up":
			a -= distance
		case "down":
			a += distance
		default:
			fmt.Println("THIS SHOULD NOT HAPPEN")
		}

	}

	fmt.Printf("Range * Depth is %d\n", r*d)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
