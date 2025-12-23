#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 1 ]]; then
  echo "usage: $0 YYDD" >&2
  exit 1
fi

day="$1"
if [[ ! "$day" =~ ^[0-9]{4}$ ]]; then
  echo "error: day must be a 4-digit YYDD value" >&2
  exit 1
fi

yy="${day:0:2}"
dd="${day:2:2}"
year="20${yy}"

cmd_file="cmd/${day}.go"
if [[ -e "$cmd_file" ]]; then
  echo "error: ${cmd_file} already exists" >&2
  exit 1
fi

cat <<EOF2 > "$cmd_file"
package cmd

import (
	"fmt"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day${day}Cmd = &cobra.Command{
	Use:                   "day${day}",
	Short:                 "${year} Advent of Code Day ${dd}",
	DisableFlagsInUseLine: true,
	Run:                   day${day}Func,
}

func init() {
	rootCmd.AddCommand(day${day}Cmd)
}

func day${day}Func(cmd *cobra.Command, args []string) {
	file, err := utils.Opener("data/${day}.txt", true)
	if err != nil {
		fmt.Printf("failed to open input: %v\n", err)
		return
	}

	_ = file

	fmt.Println("TODO: implement day ${day}")
}
EOF2

echo "created ${cmd_file}"
