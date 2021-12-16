package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var day2101Cmd = &cobra.Command{
	Use:                   "day2101",
	Short:                 "2021 Advent of Code Day 1",
	DisableFlagsInUseLine: true,
	Run:                   day2101Func,
}

func init() {
	rootCmd.AddCommand(day2101Cmd)
}

func day2101Func(cmd *cobra.Command, args []string) {

	path, _ := os.Executable()
	fmt.Println(path)

	file, err := os.Open("data/2101.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var a, b, totalDeeper, i int

	for scanner.Scan() {
		b, _ = strconv.Atoi(scanner.Text())

		// skip the first line, total all the deeper spots
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
