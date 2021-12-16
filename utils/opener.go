package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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