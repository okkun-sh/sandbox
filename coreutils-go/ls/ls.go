package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}

func run() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	output, err := list(dir)
	if err != nil {
		return err
	}

	fmt.Print(output)

	return nil
}

func list(dir string) (string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	filenames := make([]string, len(files))
	for i, file := range files {
		filenames[i] = file.Name()
	}

	maxLen := 0
	for _, name := range filenames {
		if maxLen < len(name) {
			maxLen = len(name)
		}
	}

	output := ""
	for _, name := range filenames {
		padding := maxLen + 1 - len(name)
		output += fmt.Sprintf("%s%*s", name, padding, " ")
	}
	output = strings.TrimSuffix(output, " ")

	return output, nil
}
