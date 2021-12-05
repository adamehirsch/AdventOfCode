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

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var i int // index
	var gamma int
	var epsilon int

	fc := make(map[int]int) // frequency count accumulator

	for scanner.Scan() {

		s := strings.Split(scanner.Text(), "")

		// for each index position on each line, count the 1
		for m := 0; m < len(s); m++ {
			n, _ := strconv.Atoi(s[m])
			if n == 1 {
				fc[m]++
			}
		}
		i++
	}

	// walk the frequency counter
	for c := 0; c < len(fc); c++ {

		// if more than half of the digits at that index were 1, add 2 to the power of MATH
		// ... the MATH part involved an off by one error around the length of the frequency counter

		a := int(math.Pow(2, float64(len(fc)-c-1)))

		if fc[c] > i/2 {
			gamma += a
		} else {
			epsilon += a
		}
	}

	fmt.Printf("gamma:   %13b\n", gamma)
	fmt.Printf("epsilon: %13b\n", epsilon)

	fmt.Printf("%b * %b = %d", gamma, epsilon, gamma*epsilon)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
