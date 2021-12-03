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

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var i int // index
	var gamma []int
	var epsilon []int

	fc := make(map[int]int) // frequency count accumulator

	for scanner.Scan() {

		s := strings.Split(scanner.Text(), "")

		for m := 0; m < len(s); m++ {
			n, _ := strconv.Atoi(s[m])
			if n == 1 {
				fc[m]++
			}
		}
		i++
	}

	fmt.Printf("%v: %d\n", fc, i)

	// walk the frequency accumulator
	for c := 0; c < len(fc); c++ {
		if fc[c] > i/2 {
			gamma = append(gamma, 1)
			epsilon = append(epsilon, 0)
		} else {
			gamma = append(gamma, 0)
			epsilon = append(epsilon, 1)
		}
	}

	fmt.Println(gamma)
	fmt.Println(epsilon)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
