package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type Point struct {
	X int
	Y int
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func Reverse(s []string) []string {
	r := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}

func Contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func Opener(f string, isData bool) (string, error) {

	if isData {
		pwd, _ := os.Getwd()
		f = filepath.Join(pwd, f)
	}

	file, err := os.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.TrimSuffix(string(file), "\n"), nil
}
